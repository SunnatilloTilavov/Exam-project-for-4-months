
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

type scheduleRepo struct {
    db *pgxpool.Pool
}

func NewScheduleRepo(db *pgxpool.Pool) *scheduleRepo {
    return &scheduleRepo{
        db: db,
    }
}

func (r *scheduleRepo) Create(ctx context.Context, req *ct.CreateScheduleRequest) (*ct.ScheduleResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO schedule (journalId, date, startTime, endTime, lesson, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `, req.JournalId, req.Date, req.StartTime, req.EndTime, req.Lesson, createdAt).Scan(&id)
    if err != nil {
        log.Println("error while creating schedule:", err)
        return nil, err
    }

    return &ct.ScheduleResponse{
        Id:        id.String(),
        JournalId: req.JournalId,
        Date:      req.Date,
        StartTime: req.StartTime,
        EndTime:   req.EndTime,
        Lesson:    req.Lesson,
        CreatedAt: createdAt,
        UpdatedAt: "",
        DeletedAt: "",
    }, nil
}

func (r *scheduleRepo) GetByID(ctx context.Context, req *ct.ScheduleID) (*ct.GetScheduleResponse, error) {
    var resp ct.GetScheduleResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT journalId, date, startTime, endTime, lesson, created_at, updated_at, deleted_at
        FROM schedule
        WHERE id = $1
    `, req.Id).Scan(
        &resp.JournalId, &resp.Date, &resp.StartTime, &resp.EndTime, &resp.Lesson,
        &created_at, &updated_at, &deleted_at,
    )
    resp.Id = req.Id
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting schedule by id:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *scheduleRepo) GetList(ctx context.Context, req *ct.GetListScheduleRequest) (*ct.GetListScheduleResponse, error) {
    var resp ct.GetListScheduleResponse
    var created_at, updated_at, deleted_at sql.NullString

    // Calculate OFFSET based on page and limit
    offset := (req.Page - 1) * req.Limit

    // Initialize filter string
    filter := ""

    // Add search condition if req.Search is not empty
    if req.Search != "" {
        filter = fmt.Sprintf(`
            WHERE (lesson ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search)
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    // Construct the final query with filter, limit, and offset
    query := fmt.Sprintf(`
        SELECT id, journalId, date, startTime, endTime, lesson, created_at, updated_at, deleted_at
        FROM schedule
        %s
        LIMIT $1 OFFSET $2
    `, filter)

    // Execute the query
    rows, err := r.db.Query(ctx, query, req.Limit, offset)
    if err != nil {
        log.Println("error while getting schedule list:", err)
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate response
    for rows.Next() {
        var schedule ct.GetScheduleResponse
        err := rows.Scan(&schedule.Id, &schedule.JournalId, &schedule.Date, &schedule.StartTime, &schedule.EndTime, &schedule.Lesson,
            &created_at, &updated_at, &deleted_at)
        if err != nil {
            log.Println("error while scanning schedule row:", err)
            return nil, err
        }
        schedule.CreatedAt = created_at.String
        schedule.UpdatedAt = updated_at.String
        schedule.DeletedAt = deleted_at.String
        resp.Schedules = append(resp.Schedules, &schedule)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over schedule rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *scheduleRepo) Update(ctx context.Context, req *ct.UpdateScheduleRequest) (*ct.GetScheduleResponse, error) {
    updated_at := time.Now().Format("2006-01-02 15:04:05")

    // Perform the update operation
    _, err := r.db.Exec(ctx, `
        UPDATE schedule
        SET journalId = $1, date = $2, startTime = $3, endTime = $4, lesson = $5, updated_at = $6
        WHERE id = $7
    `, req.JournalId, req.Date, req.StartTime, req.EndTime, req.Lesson, updated_at, req.Id)

    if err != nil {
        log.Println("error while updating schedule:", err)
        return nil, err
    }

    schedule, err := r.GetByID(ctx, &ct.ScheduleID{Id: req.Id})
    if err != nil {
        log.Println("error while getting schedule by id after update:", err)
        return nil, err
    }

    return schedule, nil
}

func (r *scheduleRepo) Delete(ctx context.Context, req *ct.ScheduleID) (*ct.ScheduleEmpty, error) {
    resp := &ct.ScheduleEmpty{}
    deleted_at := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE schedule
        SET deleted_at = $1
        WHERE id = $2
    `, deleted_at, req.Id)

    if err != nil {
        log.Println("error while deleting schedule:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}
