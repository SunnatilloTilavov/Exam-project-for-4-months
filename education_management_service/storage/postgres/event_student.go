package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	ct "education_management_service/genproto/education_management_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type eventStudentRepo struct {
	db *pgxpool.Pool
}

func NewEventStudentRepo(db *pgxpool.Pool) *eventStudentRepo {
	return &eventStudentRepo{
		db: db,
	}
}

func (r *eventStudentRepo) Create(ctx context.Context, req *ct.CreateEventStudentRequest) (*ct.EventStudentResponse, error) {
	var id uuid.UUID
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	err := r.db.QueryRow(ctx, `
		INSERT INTO event_student (eventId, studentId, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`, req.EventId, req.StudentId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating event_student:", err)
		return nil, err
	}

	return &ct.EventStudentResponse{
		Id:        id.String(),
		EventId:   req.EventId,
		StudentId: req.StudentId,
		CreatedAt: createdAt,
		UpdatedAt: "",
		DeletedAt: "",
	}, nil
}

func (r *eventStudentRepo) GetByID(ctx context.Context, req *ct.EventStudentID) (*ct.GetEventStudentResponse, error) {
	var resp ct.GetEventStudentResponse
	var created_at, updated_at, deleted_at sql.NullString

	err := r.db.QueryRow(ctx, `
		SELECT eventId, studentId, created_at, updated_at, deleted_at
		FROM event_student
		WHERE id = $1
	`, req.Id).Scan(
		&resp.EventId, &resp.StudentId, &created_at, &updated_at, &deleted_at,
	)
	resp.Id = req.Id
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting event_student by id:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *eventStudentRepo) GetList(ctx context.Context, req *ct.GetListEventStudentRequest) (*ct.GetListEventStudentResponse, error) {
	var resp ct.GetListEventStudentResponse
	var created_at, updated_at, deleted_at sql.NullString

	offset := (req.Page - 1) * req.Limit
	filter := ""

	if req.Search != "" {
		filter = fmt.Sprintf(`
			WHERE (eventId ILIKE '%%%v%%' OR studentId ILIKE '%%%v%%') AND deleted_at IS NULL
		`, req.Search, req.Search)
	} else {
		filter = " WHERE deleted_at IS NULL "
	}

	query := fmt.Sprintf(`
		SELECT id, eventId, studentId, created_at, updated_at, deleted_at
		FROM event_student
		%s
		LIMIT $1 OFFSET $2
	`, filter)

	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting event_student list:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eventStudent ct.GetEventStudentResponse
		err := rows.Scan(&eventStudent.Id, &eventStudent.EventId, &eventStudent.StudentId,
			&created_at, &updated_at, &deleted_at)
		if err != nil {
			log.Println("error while scanning event_student row:", err)
			return nil, err
		}
		eventStudent.CreatedAt = created_at.String
		eventStudent.UpdatedAt = updated_at.String
		eventStudent.DeletedAt = deleted_at.String
		resp.EventStudents = append(resp.EventStudents, &eventStudent)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over event_student rows:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *eventStudentRepo) Update(ctx context.Context, req *ct.UpdateEventStudentRequest) (*ct.GetEventStudentResponse, error) {
	updated_at := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
		UPDATE event_student
		SET eventId = $1, studentId = $2, updated_at = $3
		WHERE id = $4
	`, req.EventId, req.StudentId, updated_at, req.Id)

	if err != nil {
		log.Println("error while updating event_student:", err)
		return nil, err
	}

	eventStudent, err := r.GetByID(ctx, &ct.EventStudentID{Id: req.Id})
	if err != nil {
		log.Println("error while getting event_student by id after update:", err)
		return nil, err
	}

	return eventStudent, nil
}

func (r *eventStudentRepo) Delete(ctx context.Context, req *ct.EventStudentID) (*ct.EventStudentEmpty, error) {
	resp := &ct.EventStudentEmpty{}
	deleted_at := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
		UPDATE event_student
		SET deleted_at = $1
		WHERE id = $2
	`, deleted_at, req.Id)

	if err != nil {
		log.Println("error while deleting event_student:", err)
		return nil, err
	}

	resp.Msg = "Successful"
	return resp, nil
}

// func (r *eventStudentRepo) GetStudentWithEventsByID(ctx context.Context, req *ct.StudentID) (*ct.GetStudentWithEventsResponse, error) {
// 	var studentResp ct.GetStudentWithEventsResponse

// 	// Query to join student, event_student, and event tables
// 	rows, err := r.db.Query(ctx, `
// 		SELECT s.id, s.name, s.email, s.phone,
// 		       es.id as event_student_id, es.eventId, e.assignStudent, e.topic, e.startTime, e.date, e.branchId, e.created_at as event_created_at, e.updated_at as event_updated_at, e.deleted_at as event_deleted_at
// 		FROM student s
// 		LEFT JOIN event_student es ON s.id = es.studentId
// 		LEFT JOIN event e ON es.eventId = e.id
// 		WHERE s.id = $1
// 	`, req.Id)
// 	if err != nil {
// 		log.Println("error while getting student with events:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var eventDetail ct.EventDetails
// 		var studentID, name, email, phone string
// 		var eventStudentID, eventID, assignStudent, topic, startTime, date, branchID, eventCreatedAt, eventUpdatedAt, eventDeletedAt sql.NullString

// 		err := rows.Scan(&studentID, &name, &email, &phone,
// 			&eventStudentID, &eventID, &assignStudent, &topic, &startTime, &date, &branchID, &eventCreatedAt, &eventUpdatedAt, &eventDeletedAt)
// 		if err != nil {
// 			log.Println("error while scanning student with events row:", err)
// 			return nil, err
// 		}

// 		// Set student fields only once (for the first row)
// 		if studentResp.Id == "" {
// 			studentResp.Id = studentID
// 			studentResp.Name = name
// 			studentResp.Email = email
// 			studentResp.Phone = phone
// 		}

// 		// Append event details to events slice
// 		if eventID.String != "" {
// 			eventDetail.Id = eventID.String
// 			eventDetail.AssignStudent = assignStudent.String
// 			eventDetail.Topic = topic.String
// 			eventDetail.StartTime = startTime.String
// 			eventDetail.Date = date.String
// 			eventDetail.BranchId = branchID.String
// 			eventDetail.CreatedAt = eventCreatedAt.String
// 			eventDetail.UpdatedAt = eventUpdatedAt.String
// 			eventDetail.DeletedAt = eventDeletedAt.String

// 			studentResp.Events = append(studentResp.Events, &eventDetail)
// 		}
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Println("error after iterating over student with events rows:", err)
// 		return nil, err
// 	}

// 	return &studentResp, nil
// }

func (r *eventStudentRepo) GetStudentWithEventsByID(ctx context.Context, req *ct.StudentID) (*ct.GetStudentWithEventsResponse, error) {
    var studentResp ct.GetStudentWithEventsResponse

    // Query to join student, event_student, and event tables
    rows, err := r.db.Query(ctx, `
        SELECT s.id, s.name, s.email, s.phone,
               es.id as event_student_id, es.eventId, e.assignStudent, e.topic, e.startTime, e.date, e.branchId, e.created_at as event_created_at, e.updated_at as event_updated_at, e.deleted_at as event_deleted_at
        FROM student s
        LEFT JOIN event_student es ON s.id = es.studentId
        LEFT JOIN event e ON es.eventId = e.id
        WHERE s.id = $1
    `, req.Id)
    if err != nil {
        log.Println("error while getting student with events:", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var studentID, name, email, phone string
        var eventStudentID, eventID, assignStudent, topic, startTime, date, branchID, eventCreatedAt, eventUpdatedAt, eventDeletedAt sql.NullString

        err := rows.Scan(&studentID, &name, &email, &phone,
            &eventStudentID, &eventID, &assignStudent, &topic, &startTime, &date, &branchID, &eventCreatedAt, &eventUpdatedAt, &eventDeletedAt)
        if err != nil {
            log.Println("error while scanning student with events row:", err)
            return nil, err
        }

        // Set student fields only once (for the first row)
        if studentResp.Id == "" {
            studentResp.Id = studentID
            studentResp.Name = name
            studentResp.Email = email
            studentResp.Phone = phone
        }

        // Create an EventStudentResponse instance and populate it
        eventResponse := &ct.EventStudentResponse{
            Id:           eventID.String,
            EventId:      eventID.String,
            StudentId:    studentID, // Assuming student ID is required in EventStudentResponse
            CreatedAt:    eventCreatedAt.String,
            UpdatedAt:    eventUpdatedAt.String,
            DeletedAt:    eventDeletedAt.String,
            AssignStudent: assignStudent.String,
            Topic:        topic.String,
            StartTime:    startTime.String,
            Date:         date.String,
            BranchId:     branchID.String,
        }

        // Append the EventStudentResponse instance to studentResp.Events
        studentResp.Events = append(studentResp.Events, eventResponse)
    }

    if err := rows.Err(); err != nil {
        log.Println("error after iterating over student with events rows:", err)
        return nil, err
    }

    return &studentResp, nil
}
