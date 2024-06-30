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

type jurnalRepo struct {
    db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewJurnalRepo(db *pgxpool.Pool) *jurnalRepo {
    return &jurnalRepo{
        db: db,
    }
}

func (r *jurnalRepo) Create(ctx context.Context, req *ct.CreateJurnalRequest) (*ct.JurnalResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO journal (groupId, fromDate, toDate, studentsCount, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `, req.GroupId, req.FromDate, req.ToDate, req.StudentsCount, createdAt).Scan(&id)
    if err != nil {
        log.Println("error while creating journal:", err)
        return nil, err
    }

    return &ct.JurnalResponse{
        Id:            id.String(),
        GroupId:       req.GroupId,
        FromDate:      req.FromDate,
        ToDate:        req.ToDate,
        StudentsCount: req.StudentsCount,
        CreatedAt:     createdAt,
        UpdatedAt:     "",
        DeletedAt:     "",
    }, nil
}

func (r *jurnalRepo) GetByID(ctx context.Context, req *ct.JurnalID) (*ct.GetJurnalResponse, error) {
    var resp ct.GetJurnalResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT groupId, fromDate, toDate, studentsCount, created_at, updated_at, deleted_at
        FROM journal
        WHERE id = $1
    `, req.Id).Scan(
        &resp.GroupId, &resp.FromDate, &resp.ToDate, &resp.StudentsCount,
        &created_at, &updated_at, &deleted_at,
    )
    resp.Id = req.Id
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting journal by id:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *jurnalRepo) GetList(ctx context.Context, req *ct.GetListJurnalRequest) (*ct.GetListJurnalResponse, error) {
    var resp ct.GetListJurnalResponse
    var created_at, updated_at, deleted_at sql.NullString

    // Calculate OFFSET based on page and limit
    offset := (req.Page - 1) * req.Limit

    // Initialize filter string
    filter := ""

    // Add search condition if req.Search is not empty
    if req.Search != "" {
        filter = fmt.Sprintf(`
            WHERE (groupId ILIKE '%%%v%%' OR fromDate::text ILIKE '%%%v%%' OR toDate::text ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search, req.Search, req.Search)
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    // Construct the final query with filter, limit, and offset
    query := fmt.Sprintf(`
        SELECT id, groupId, fromDate, toDate, studentsCount, created_at, updated_at, deleted_at
        FROM journal
        %s
        LIMIT $1 OFFSET $2
    `, filter)

    // Execute the query
    rows, err := r.db.Query(ctx, query, req.Limit, offset)
    if err != nil {
        log.Println("error while getting journal list:", err)
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate response
    for rows.Next() {
        var journal ct.GetJurnalResponse
        err := rows.Scan(&journal.Id, &journal.GroupId, &journal.FromDate, &journal.ToDate, &journal.StudentsCount,
            &created_at, &updated_at, &deleted_at)
        if err != nil {
            log.Println("error while scanning journal row:", err)
            return nil, err
        }
        journal.CreatedAt = created_at.String
        journal.UpdatedAt = updated_at.String
        journal.DeletedAt = deleted_at.String
        resp.Jurnals = append(resp.Jurnals, &journal)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over journal rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *jurnalRepo) Update(ctx context.Context, req *ct.UpdateJurnalRequest) (*ct.GetJurnalResponse, error) {
    updated_at := time.Now().Format("2006-01-02 15:04:05")

    // Perform the update operation
    _, err := r.db.Exec(ctx, `
        UPDATE journal
        SET groupId = $1, fromDate = $2, toDate = $3, studentsCount = $4, updated_at = $5
        WHERE id = $6
    `, req.GroupId, req.FromDate, req.ToDate, req.StudentsCount, updated_at, req.Id)

    if err != nil {
        log.Println("error while updating journal:", err)
        return nil, err
    }

    journal, err := r.GetByID(ctx, &ct.JurnalID{Id: req.Id})
    if err != nil {
        log.Println("error while getting journal by id after update:", err)
        return nil, err
    }

    return journal, nil
}

func (r *jurnalRepo) Delete(ctx context.Context, req *ct.JurnalID) (*ct.JurnalEmpty, error) {
    resp := &ct.JurnalEmpty{}
    deleted_at := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE journal
        SET deleted_at = $1
        WHERE id = $2
    `, deleted_at, req.Id)

    if err != nil {
        log.Println("error while deleting journal:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}



func (r *jurnalRepo) GetByGroupID(ctx context.Context, req *ct.GroupId) (*ct.GetJurnalResponse, error) {
    var resp ct.GetJurnalResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT id,groupId, fromDate, toDate, studentsCount, created_at, updated_at, deleted_at
        FROM journal
        WHERE groupId = $1
    `, req.Id).Scan(
        &resp.Id,&resp.GroupId, &resp.FromDate, &resp.ToDate, &resp.StudentsCount,
        &created_at, &updated_at, &deleted_at,
    )
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting journal by id:", err)
        return nil, err
    }
    return &resp, nil
}