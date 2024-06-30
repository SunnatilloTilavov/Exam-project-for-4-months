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

type branchRepo struct {
    db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
    return &branchRepo{
        db: db,
    }
}

func (r *branchRepo) Create(ctx context.Context, req *ct.CreateBranchRequest) (*ct.BranchResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO branch (name, address, phone, created_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, req.Name, req.Address, req.Phone, createdAt).Scan(&id)
    if err != nil {
        log.Println("error while creating branch:", err)
        return nil, err
    }

    return &ct.BranchResponse{
        Id:        id.String(),
        Name:      req.Name,
        Address:   req.Address,
        Phone:     req.Phone,
        CreatedAt: createdAt,
        UpdatedAt: "",
        DeletedAt: "",
    }, nil
}

func (r *branchRepo) GetByID(ctx context.Context, req *ct.BranchID) (*ct.GetBranchResponse, error) {
    var resp ct.GetBranchResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT name, address, phone, created_at, updated_at, deleted_at
        FROM branch
        WHERE id = $1
    `, req.Id).Scan(
        &resp.Name, &resp.Address, &resp.Phone,
        &created_at, &updated_at, &deleted_at,
    )
    resp.Id = req.Id
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting branch by id:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *branchRepo) GetList(ctx context.Context, req *ct.GetListBranchRequest) (*ct.GetListBranchResponse, error) {
    var resp ct.GetListBranchResponse
    var created_at, updated_at, deleted_at sql.NullString

    // Calculate OFFSET based on page and limit
    offset := (req.Page - 1) * req.Limit

    // Initialize filter string
    filter := ""

    // Add search condition if req.Search is not empty
    if req.Search != "" {
        filter = fmt.Sprintf(`
            WHERE (name ILIKE '%%%v%%'
            OR address ILIKE '%%%v%%'
            OR phone ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search, req.Search, req.Search)
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    // Construct the final query with filter, limit, and offset
    query := fmt.Sprintf(`
        SELECT id, name, address, phone, created_at, updated_at, deleted_at
        FROM branch
        %s
        LIMIT $1 OFFSET $2
    `, filter)

    // Execute the query
    rows, err := r.db.Query(ctx, query, req.Limit, offset)
    if err != nil {
        log.Println("error while getting branch list:", err)
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate response
    for rows.Next() {
        var branch ct.GetBranchResponse
        err := rows.Scan(&branch.Id, &branch.Name, &branch.Address, &branch.Phone,
            &created_at, &updated_at, &deleted_at)
        if err != nil {
            log.Println("error while scanning branch row:", err)
            return nil, err
        }
        branch.CreatedAt = created_at.String
        branch.UpdatedAt = updated_at.String
        branch.DeletedAt = deleted_at.String
        resp.Branches = append(resp.Branches, &branch)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over branch rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *branchRepo) Update(ctx context.Context, req *ct.UpdateBranchRequest) (*ct.GetBranchResponse, error) {
    updated_at := time.Now().Format("2006-01-02 15:04:05")

    // Perform the update operation
    _, err := r.db.Exec(ctx, `
        UPDATE branch
        SET name = $1, address = $2, phone = $3, updated_at = $4
        WHERE id = $5
    `, req.Name, req.Address, req.Phone, updated_at, req.Id)

    if err != nil {
        log.Println("error while updating branch:", err)
        return nil, err
    }

    branch, err := r.GetByID(ctx, &ct.BranchID{Id: req.Id})
    if err != nil {
        log.Println("error while getting branch by id after update")
        return nil, err
    }

    return branch, nil
}

func (r *branchRepo) Delete(ctx context.Context, req *ct.BranchID) (*ct.BranchEmpty, error) {
    resp := &ct.BranchEmpty{}
    deleted_at := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE branch
        SET deleted_at = $1
        WHERE id = $2
    `, deleted_at, req.Id)

    if err != nil {
        log.Println("error while deleting branch:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}
