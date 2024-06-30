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

type taskRepo struct {
    db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewTaskRepo(db *pgxpool.Pool) *taskRepo {
    return &taskRepo{
        db: db,
    }
}

func (r *taskRepo) Create(ctx context.Context, req *ct.CreateTaskRequest) (*ct.TaskResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO task (scheduleId, label, deadlineDate, deadlineTime, score, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `, req.ScheduleId, req.Label, req.DeadlineDate, req.DeadlineTime, req.Score, createdAt).Scan(&id)
    if err != nil {
        log.Println("error while creating task:", err)
        return nil, err
    }

    return &ct.TaskResponse{
        Id:           id.String(),
        ScheduleId:   req.ScheduleId,
        Label:        req.Label,
        DeadlineDate: req.DeadlineDate,
        DeadlineTime: req.DeadlineTime,
        Score:        req.Score,
        CreatedAt:    createdAt,
        UpdatedAt:    "",
        DeletedAt:    "",
    }, nil
}

func (r *taskRepo) GetByID(ctx context.Context, req *ct.TaskID) (*ct.GetTaskResponse, error) {
    var resp ct.GetTaskResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT scheduleId, label, deadlineDate, deadlineTime, score, created_at, updated_at, deleted_at
        FROM task
        WHERE id = $1
    `, req.Id).Scan(
        &resp.ScheduleId, &resp.Label, &resp.DeadlineDate, &resp.DeadlineTime, &resp.Score,
        &created_at, &updated_at, &deleted_at,
    )
    resp.Id = req.Id
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting task by id:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *taskRepo) GetList(ctx context.Context, req *ct.GetListTaskRequest) (*ct.GetListTaskResponse, error) {
    var resp ct.GetListTaskResponse
    var created_at, updated_at, deleted_at sql.NullString

    // Calculate OFFSET based on page and limit
    offset := (req.Page - 1) * req.Limit

    // Initialize filter string
    filter := ""

    // Add search condition if req.Search is not empty
    if req.Search != "" {
        filter = fmt.Sprintf(`
            WHERE (label ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search)
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    // Construct the final query with filter, limit, and offset
    query := fmt.Sprintf(`
        SELECT id, scheduleId, label, deadlineDate, deadlineTime, score, created_at, updated_at, deleted_at
        FROM task
        %s
        LIMIT $1 OFFSET $2
    `, filter)

    // Execute the query
    rows, err := r.db.Query(ctx, query, req.Limit, offset)
    if err != nil {
        log.Println("error while getting task list:", err)
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate response
    for rows.Next() {
        var task ct.GetTaskResponse
        err := rows.Scan(&task.Id, &task.ScheduleId, &task.Label, &task.DeadlineDate, &task.DeadlineTime, &task.Score,
            &created_at, &updated_at, &deleted_at)
        if err != nil {
            log.Println("error while scanning task row:", err)
            return nil, err
        }
        task.CreatedAt = created_at.String
        task.UpdatedAt = updated_at.String
        task.DeletedAt = deleted_at.String
        resp.Tasks = append(resp.Tasks, &task)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over task rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *taskRepo) Update(ctx context.Context, req *ct.UpdateTaskRequest) (*ct.GetTaskResponse, error) {
    updated_at := time.Now().Format("2006-01-02 15:04:05")

    // Perform the update operation
    _, err := r.db.Exec(ctx, `
        UPDATE task
        SET scheduleId = $1, label = $2, deadlineDate = $3, deadlineTime = $4, score = $5, updated_at = $6
        WHERE id = $7
    `, req.ScheduleId, req.Label, req.DeadlineDate, req.DeadlineTime, req.Score, updated_at, req.Id)

    if err != nil {
        log.Println("error while updating task:", err)
        return nil, err
    }

    task, err := r.GetByID(ctx, &ct.TaskID{Id: req.Id})
    if err != nil {
        log.Println("error while getting task by id after update:", err)
        return nil, err
    }

    return task, nil
}

func (r *taskRepo) Delete(ctx context.Context, req *ct.TaskID) (*ct.TaskEmpty, error) {
    resp := &ct.TaskEmpty{}
    deleted_at := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE task
        SET deleted_at = $1
        WHERE id = $2
    `, deleted_at, req.Id)

    if err != nil {
        log.Println("error while deleting task:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}
