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

type studentPaymentRepo struct {
    db *pgxpool.Pool
}

func NewStudentPaymentRepo(db *pgxpool.Pool) *studentPaymentRepo {
    return &studentPaymentRepo{
        db: db,
    }
}

func (r *studentPaymentRepo) Create(ctx context.Context, req *ct.CreateStudentPaymentRequest) (*ct.StudentPaymentResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO student_payment (studentId, groupId, paidSum, administrationId, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `, req.StudentId, req.GroupId, req.PaidSum, req.AdministrationId, createdAt).Scan(&id)
    if err != nil {
        log.Println("error while creating student_payment:", err)
        return nil, err
    }

    return &ct.StudentPaymentResponse{
        Id:              id.String(),
        StudentId:       req.StudentId,
        GroupId:         req.GroupId,
        PaidSum:         req.PaidSum,
        AdministrationId: req.AdministrationId,
        CreatedAt:       createdAt,
    }, nil
}

func (r *studentPaymentRepo) GetByID(ctx context.Context, req *ct.StudentPaymentID) (*ct.GetStudentPaymentResponse, error) {
    var resp ct.GetStudentPaymentResponse
    var createdAt, updatedAt, deletedAt sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT studentId, groupId, paidSum, administrationId, created_at, updated_at, deleted_at
        FROM student_payment
        WHERE id = $1
    `, req.Id).Scan(&resp.StudentId, &resp.GroupId, &resp.PaidSum, &resp.AdministrationId, &createdAt, &updatedAt, &deletedAt)
    if err != nil {
        log.Println("error while getting student_payment by id:", err)
        return nil, err
    }
    resp.Id = req.Id
    resp.CreatedAt = createdAt.String
    resp.UpdatedAt = updatedAt.String
    resp.DeletedAt = deletedAt.String
    return &resp, nil
}

func (r *studentPaymentRepo) GetList(ctx context.Context, req *ct.GetListStudentPaymentRequest) (*ct.GetListStudentPaymentResponse, error) {
    var resp ct.GetListStudentPaymentResponse
    var createdAt, updatedAt, deletedAt sql.NullString

    offset := (req.Page - 1) * req.Limit
    filter := ""
    if req.Search != "" {
        filter = `
            WHERE (studentId ILIKE '%' || $3 || '%' OR groupId ILIKE '%' || $3 || '%' OR administrationId ILIKE '%' || $3 || '%') AND deleted_at IS NULL
        `
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    query := `
        SELECT id, studentId, groupId, paidSum, administrationId, created_at, updated_at, deleted_at
        FROM student_payment
        ` + filter + `
        LIMIT $1 OFFSET $2
    `
    rows, err := r.db.Query(ctx, query, req.Limit, offset, req.Search)
    if err != nil {
        log.Println("error while getting student_payment list:", err)
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var studentPayment ct.GetStudentPaymentResponse
        err := rows.Scan(&studentPayment.Id, &studentPayment.StudentId, &studentPayment.GroupId, &studentPayment.PaidSum, &studentPayment.AdministrationId, &createdAt, &updatedAt, &deletedAt)
        if err != nil {
            log.Println("error while scanning student_payment row:", err)
            return nil, err
        }
        studentPayment.CreatedAt = createdAt.String
        studentPayment.UpdatedAt = updatedAt.String
        studentPayment.DeletedAt = deletedAt.String
        resp.StudentPayments = append(resp.StudentPayments, &studentPayment)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over student_payment rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *studentPaymentRepo) Update(ctx context.Context, req *ct.UpdateStudentPaymentRequest) (*ct.GetStudentPaymentResponse, error) {
    updatedAt := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE student_payment
        SET studentId = $1, groupId = $2, paidSum = $3, administrationId = $4, updated_at = $5
        WHERE id = $6
    `, req.StudentId, req.GroupId, req.PaidSum, req.AdministrationId, updatedAt, req.Id)
    if err != nil {
        log.Println("error while updating student_payment:", err)
        return nil, err
    }

    studentPayment, err := r.GetByID(ctx, &ct.StudentPaymentID{Id: req.Id})
    if err != nil {
        log.Println("error while getting student_payment by id after update:", err)
        return nil, err
    }

    return studentPayment, nil
}

func (r *studentPaymentRepo) Delete(ctx context.Context, req *ct.StudentPaymentID) (*ct.StudentPaymentEmpty, error) {
    resp := &ct.StudentPaymentEmpty{}
    deletedAt := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE student_payment
        SET deleted_at = $1
        WHERE id = $2
    `, deletedAt, req.Id)
    if err != nil {
        log.Println("error while deleting student_payment:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}
