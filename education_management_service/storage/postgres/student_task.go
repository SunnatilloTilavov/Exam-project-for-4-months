package postgres

import (
    "context"
    "database/sql"
    "log"
    "time"

    ct "education_management_service/genproto/education_management_service"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v4/pgxpool"
)

type studentTaskRepo struct {
    db *pgxpool.Pool
}

func NewStudentTaskRepo(db *pgxpool.Pool) *studentTaskRepo {
    return &studentTaskRepo{
        db: db,
    }
}

func (r *studentTaskRepo) Create(ctx context.Context, req *ct.CreateStudentTaskRequest) (*ct.StudentTaskResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO student_task (taskId, studentId, created_at, score)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, req.TaskId, req.StudentId, createdAt, req.Score).Scan(&id)
    if err != nil {
        log.Println("error while creating student_task:", err)
        return nil, err
    }

    return &ct.StudentTaskResponse{
        Id:        id.String(),
        TaskId:    req.TaskId,
        StudentId: req.StudentId,
        CreatedAt: createdAt,
        Score:     req.Score,
    }, nil
}

func (r *studentTaskRepo) GetByID(ctx context.Context, req *ct.StudentTaskID) (*ct.GetStudentTaskResponse, error) {
    var resp ct.GetStudentTaskResponse
    var createdAt, updatedAt, deletedAt sql.NullString
    var score sql.NullInt32

    err := r.db.QueryRow(ctx, `
        SELECT taskId, studentId, created_at, updated_at, deleted_at, score
        FROM student_task
        WHERE id = $1
    `, req.Id).Scan(&resp.TaskId, &resp.StudentId, &createdAt, &updatedAt, &deletedAt, &score)
    if err != nil {
        log.Println("error while getting student_task by id:", err)
        return nil, err
    }
    resp.Id = req.Id
    resp.CreatedAt = createdAt.String
    resp.UpdatedAt = updatedAt.String
    resp.DeletedAt = deletedAt.String
    if score.Valid {
        resp.Score = score.Int32
    }
    return &resp, nil
}

func (r *studentTaskRepo) GetList(ctx context.Context, req *ct.GetListStudentTaskRequest) (*ct.GetListStudentTaskResponse, error) {
    var resp ct.GetListStudentTaskResponse
    var createdAt, updatedAt, deletedAt sql.NullString
    var score sql.NullInt32

    offset := (req.Page - 1) * req.Limit
    filter := ""
    if req.Search != "" {
        filter = `
            WHERE (taskId ILIKE '%' || $3 || '%' OR studentId ILIKE '%' || $3 || '%') AND deleted_at IS NULL
        `
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    query := `
        SELECT id, taskId, studentId, created_at, updated_at, deleted_at, score
        FROM student_task
        ` + filter + `
        LIMIT $1 OFFSET $2
    `
    rows, err := r.db.Query(ctx, query, req.Limit, offset, req.Search)
    if err != nil {
        log.Println("error while getting student_task list:", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var studentTask ct.GetStudentTaskResponse
        err := rows.Scan(&studentTask.Id, &studentTask.TaskId, &studentTask.StudentId, &createdAt, &updatedAt, &deletedAt, &score)
        if err != nil {
            log.Println("error while scanning student_task row:", err)
            return nil, err
        }
        studentTask.CreatedAt = createdAt.String
        studentTask.UpdatedAt = updatedAt.String
        studentTask.DeletedAt = deletedAt.String
        if score.Valid {
            studentTask.Score = score.Int32
        }
        resp.StudentTasks = append(resp.StudentTasks, &studentTask)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over student_task rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *studentTaskRepo) Update(ctx context.Context, req *ct.UpdateStudentTaskRequest) (*ct.GetStudentTaskResponse, error) {
    updatedAt := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE student_task
        SET taskId = $1, studentId = $2, updated_at = $3, score = $4
        WHERE id = $5
    `, req.TaskId, req.StudentId, updatedAt, req.Score, req.Id)
    if err != nil {
        log.Println("error while updating student_task:", err)
        return nil, err
    }

    studentTask, err := r.GetByID(ctx, &ct.StudentTaskID{Id: req.Id})
    if err != nil {
        log.Println("error while getting student_task by id after update:", err)
        return nil, err
    }

    return studentTask, nil
}

func (r *studentTaskRepo) Delete(ctx context.Context, req *ct.StudentTaskID) (*ct.StudentTaskEmpty, error) {
    resp := &ct.StudentTaskEmpty{}
    deletedAt := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE student_task
        SET deleted_at = $1
        WHERE id = $2
    `, deletedAt, req.Id)
    if err != nil {
        log.Println("error while deleting student_task:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}



func (r *studentTaskRepo) UpdateScoreforTeacher(ctx context.Context, req *ct.UpdateStudentScoreRequest) (*ct.GetStudentTaskResponse, error) {
    _, err := r.db.Exec(ctx, `
        UPDATE student_task
        SET score = $1
        WHERE id = $2
    `, req.Score, req.Id)
    if err != nil {
        log.Println("error while updating student_task:", err)
        return nil, err
    }

    studentTask, err := r.GetByID(ctx, &ct.StudentTaskID{Id: req.Id})
    if err != nil {
        log.Println("error while getting student_task by id after update:", err)
        return nil, err
    }

    return studentTask, nil
}


func (r *studentTaskRepo) UpdateScoreforStudent(ctx context.Context, req *ct.UpdateStudentScoreRequest) (*ct.GetStudentTaskResponse, error) {
    _, err := r.db.Exec(ctx, `
        UPDATE student_task
        SET score = $1
        WHERE id = $2
    `, req.Score, req.Id)
    if err != nil {
        log.Println("error while updating student_task:", err)
        return nil, err
    }

    studentTask, err := r.GetByID(ctx, &ct.StudentTaskID{Id: req.Id})
    if err != nil {
        log.Println("error while getting student_task by id after update:", err)
        return nil, err
    }
    return studentTask, nil
}


func (r *studentTaskRepo) GetByIDStudent(ctx context.Context, req *ct.TaskStudentID) (*ct.GetStudentTaskResponse, error) {
    var resp ct.GetStudentTaskResponse
    var createdAt, updatedAt, deletedAt sql.NullString
    var score sql.NullInt32

    err := r.db.QueryRow(ctx, `
        SELECT id,taskId, studentId, created_at, updated_at, deleted_at, score
        FROM student_task
        WHERE studentId = $1
    `, req.Id).Scan(&resp.Id,&resp.TaskId, &resp.StudentId, &createdAt, &updatedAt, &deletedAt, &score)
    if err != nil {
        log.Println("error while getting student_task by id:", err)
        return nil, err
    }
    resp.CreatedAt = createdAt.String
    resp.UpdatedAt = updatedAt.String
    resp.DeletedAt = deletedAt.String
    if score.Valid {
        resp.Score = score.Int32
    }
    return &resp, nil
}