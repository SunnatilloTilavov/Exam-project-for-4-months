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

type groupRepo struct {
    db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewGroupRepo(db *pgxpool.Pool) *groupRepo {
    return &groupRepo{
        db: db,
    }
}

func (r *groupRepo) Create(ctx context.Context, req *ct.CreateGroupRequest) (*ct.GroupResponse, error) {
    var id uuid.UUID
    createdAt := time.Now().Format("2006-01-02 15:04:05")

    err := r.db.QueryRow(ctx, `
        INSERT INTO "group" (name, teacherId, supportTeacherId, branchId, type, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `, req.Name, req.TeacherId, req.SupportTeacherId, req.BranchId, req.Type, createdAt).Scan(&id)
    if err != nil {
        log.Println("error while creating group:", err)
        return nil, err
    }

    return &ct.GroupResponse{
        Id:             id.String(),
        Name:           req.Name,
        TeacherId:      req.TeacherId,
        SupportTeacherId: req.SupportTeacherId,
        BranchId:       req.BranchId,
        Type:           req.Type,
        CreatedAt:      createdAt,
        UpdatedAt:      "",
        DeletedAt:      "",
    }, nil
}

func (r *groupRepo) GetByID(ctx context.Context, req *ct.GroupID) (*ct.GetGroupResponse, error) {
    var resp ct.GetGroupResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT name, teacherId, supportTeacherId, branchId, type, created_at, updated_at, deleted_at
        FROM "group"
        WHERE id = $1
    `, req.Id).Scan(
        &resp.Name, &resp.TeacherId, &resp.SupportTeacherId, &resp.BranchId, &resp.Type,
        &created_at, &updated_at, &deleted_at,
    )
    resp.Id = req.Id
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting group by id:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *groupRepo) GetList(ctx context.Context, req *ct.GetListGroupRequest) (*ct.GetListGroupResponse, error) {
    var resp ct.GetListGroupResponse
    var created_at, updated_at, deleted_at sql.NullString

    // Calculate OFFSET based on page and limit
    offset := (req.Page - 1) * req.Limit

    // Initialize filter string
    filter := ""

    // Add search condition if req.Search is not empty
    if req.Search != "" {
        filter = fmt.Sprintf(`
            WHERE (name ILIKE '%%%v%%' OR type ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search, req.Search)
    } else {
        filter = " WHERE deleted_at IS NULL "
    }

    // Construct the final query with filter, limit, and offset
    query := fmt.Sprintf(`
        SELECT id, name, teacherId, supportTeacherId, branchId, type, created_at, updated_at, deleted_at
        FROM "group"
        %s
        LIMIT $1 OFFSET $2
    `, filter)

    // Execute the query
    rows, err := r.db.Query(ctx, query, req.Limit, offset)
    if err != nil {
        log.Println("error while getting group list:", err)
        return nil, err
    }
    defer rows.Close()

    // Iterate over the rows and populate response
    for rows.Next() {
        var group ct.GetGroupResponse
        err := rows.Scan(&group.Id, &group.Name, &group.TeacherId, &group.SupportTeacherId, &group.BranchId, &group.Type,
            &created_at, &updated_at, &deleted_at)
        if err != nil {
            log.Println("error while scanning group row:", err)
            return nil, err
        }
        group.CreatedAt = created_at.String
        group.UpdatedAt = updated_at.String
        group.DeletedAt = deleted_at.String
        resp.Groups = append(resp.Groups, &group)
    }
    if err := rows.Err(); err != nil {
        log.Println("error after iterating over group rows:", err)
        return nil, err
    }
    return &resp, nil
}

func (r *groupRepo) Update(ctx context.Context, req *ct.UpdateGroupRequest) (*ct.GetGroupResponse, error) {
    updated_at := time.Now().Format("2006-01-02 15:04:05")

    // Perform the update operation
    _, err := r.db.Exec(ctx, `
        UPDATE "group"
        SET name = $1, teacherId = $2, supportTeacherId = $3, branchId = $4, type = $5, updated_at = $6
        WHERE id = $7
    `, req.Name, req.TeacherId, req.SupportTeacherId, req.BranchId, req.Type, updated_at, req.Id)

    if err != nil {
        log.Println("error while updating group:", err)
        return nil, err
    }

    group, err := r.GetByID(ctx, &ct.GroupID{Id: req.Id})
    if err != nil {
        log.Println("error while getting group by id after update:", err)
        return nil, err
    }

    return group, nil
}

func (r *groupRepo) Delete(ctx context.Context, req *ct.GroupID) (*ct.GroupEmpty, error) {
    resp := &ct.GroupEmpty{}
    deleted_at := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(ctx, `
        UPDATE "group"
        SET deleted_at = $1
        WHERE id = $2
    `, deleted_at, req.Id)

    if err != nil {
        log.Println("error while deleting group:", err)
        return nil, err
    }

    resp.Msg = "Successful"
    return resp, nil
}



func (r *groupRepo) GetByIDTeacher(ctx context.Context, req *ct.TeacherID) (*ct.GetGroupResponse, error) {
    var resp ct.GetGroupResponse
    var created_at, updated_at, deleted_at sql.NullString

    err := r.db.QueryRow(ctx, `
        SELECT id,name, teacherId, supportTeacherId, branchId, type, created_at, updated_at, deleted_at
        FROM "group"
        WHERE teacherId = $1
    `, req.Id).Scan(
        &resp.Id,&resp.Name, &resp.TeacherId, &resp.SupportTeacherId, &resp.BranchId, &resp.Type,
        &created_at, &updated_at, &deleted_at,
    )
    resp.CreatedAt = created_at.String
    resp.UpdatedAt = updated_at.String
    resp.DeletedAt = deleted_at.String
    if err != nil {
        log.Println("error while getting group by id:", err)
        return nil, err
    }
    return &resp, nil
}
