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

type eventRepo struct {
	db *pgxpool.Pool
}

func NewEventRepo(db *pgxpool.Pool) *eventRepo {
	return &eventRepo{
		db: db,
	}
}

func (r *eventRepo) Create(ctx context.Context, req *ct.CreateEventRequest) (*ct.EventResponse, error) {
	var id uuid.UUID
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	err := r.db.QueryRow(ctx, `
		INSERT INTO event (assignStudent, topic, startTime, date, branchId, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, req.AssignStudent, req.Topic, req.StartTime, req.Date, req.BranchId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating event:", err)
		return nil, err
	}

	return &ct.EventResponse{
		Id:             id.String(),
		AssignStudent:  req.AssignStudent,
		Topic:          req.Topic,
		StartTime:      req.StartTime,
		Date:           req.Date,
		BranchId:       req.BranchId,
		CreatedAt:      createdAt,
		UpdatedAt:      "",
		DeletedAt:      "",
	}, nil
}

func (r *eventRepo) GetByID(ctx context.Context, req *ct.EventID) (*ct.GetEventResponse, error) {
	var resp ct.GetEventResponse
	var created_at, updated_at, deleted_at,Topic,StartTime,Date sql.NullString

	err := r.db.QueryRow(ctx, `
		SELECT assignStudent, topic, startTime, date, branchId, created_at, updated_at, deleted_at
		FROM event
		WHERE id = $1
	`, req.Id).Scan(
		&resp.AssignStudent, &Topic, &StartTime, &Date, &resp.BranchId,
		&created_at, &updated_at, &deleted_at,
	)
	resp.Id = req.Id
	resp.StartTime=StartTime.String
	resp.Date=Date.String
	resp.Topic=Topic.String
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting event by id:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *eventRepo) GetList(ctx context.Context, req *ct.GetListEventRequest) (*ct.GetListEventResponse, error) {
	var resp ct.GetListEventResponse
	var created_at, updated_at, deleted_at,Topic,StartTime,Date sql.NullString

	offset := (req.Page - 1) * req.Limit

	filter := ""
	if req.Search != "" {
		filter = fmt.Sprintf(`
			WHERE (assignStudent ILIKE '%%%v%%' OR topic ILIKE '%%%v%%') AND deleted_at IS NULL
		`, req.Search, req.Search)
	} else {
		filter = " WHERE deleted_at IS NULL "
	}

	query := fmt.Sprintf(`
		SELECT id, assignStudent, topic, startTime, date, branchId, created_at, updated_at, deleted_at
		FROM event
		%s
		LIMIT $1 OFFSET $2
	`, filter)

	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting event list:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event ct.GetEventResponse
		err := rows.Scan(&event.Id, &event.AssignStudent, &Topic, &StartTime, &Date, &event.BranchId,
			&created_at, &updated_at, &deleted_at)
		if err != nil {
			log.Println("error while scanning event row:", err)
			return nil, err
		}
		event.Topic=Topic.String
		event.StartTime=StartTime.String
		event.Date=Date.String
		event.CreatedAt = created_at.String
		event.UpdatedAt = updated_at.String
		event.DeletedAt = deleted_at.String
		resp.Events = append(resp.Events, &event)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over event rows:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *eventRepo) Update(ctx context.Context, req *ct.UpdateEventRequest) (*ct.GetEventResponse, error) {
	updated_at := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
		UPDATE event
		SET assignStudent = $1, topic = $2, startTime = $3, date = $4, branchId = $5, updated_at = $6
		WHERE id = $7
	`, req.AssignStudent, req.Topic, req.StartTime, req.Date, req.BranchId, updated_at, req.Id)

	if err != nil {
		log.Println("error while updating event:", err)
		return nil, err
	}

	event, err := r.GetByID(ctx, &ct.EventID{Id: req.Id})
	if err != nil {
		log.Println("error while getting event by id after update:", err)
		return nil, err
	}

	return event, nil
}

func (r *eventRepo) Delete(ctx context.Context, req *ct.EventID) (*ct.EventEmpty, error) {
	resp := &ct.EventEmpty{}
	deleted_at := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
		UPDATE event
		SET deleted_at = $1
		WHERE id = $2
	`, deleted_at, req.Id)
	if err != nil {
		log.Println("error while deleting event:", err)
		return nil, err
	}
 resp.Msg = "Successful"
	return resp, nil
}
