package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
"strconv"
	ct "user_service/genproto/user_service"
"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type managerRepo struct {
	db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewManagerRepo(db *pgxpool.Pool) *managerRepo {
	return &managerRepo{
		db: db,
	}
}

func (r *managerRepo) Create(ctx context.Context, req *ct.CreateManagerRequest) (*ct.ManagerResponse, error) {
	var id uuid.UUID

	loginlast, err := r.GetLastLogin(ctx)
	if err != nil {
		log.Println("error while creating menager:", err)
		return nil, err
	}
	login:=GenerateNewLoginManager(loginlast)

	createdAt := time.Now().Format("2006-01-02 15:04:05")

	err = r.db.QueryRow(ctx, `
        INSERT INTO manager (login, fullname, phone, password, salary, branchId, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `, login, req.Fullname, req.Phone, req.Password, req.Salary, req.BranchId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating manager:", err)
		return nil, err
	}

	return &ct.ManagerResponse{
		Id:         id.String(),
		Login:      login,
		Fullname:   req.Fullname,
		Phone:      req.Phone,
		Password:   req.Password,
		Salary:     req.Salary,
		BranchId:   req.BranchId,
		CreatedAt:  createdAt,
		UpdatedAt:  "",
		DeletedAt:  "",
	}, nil
}

func (r *managerRepo) GetByID(ctx context.Context, req *ct.ManagerID) (*ct.GetManagerResponse, error) {
	var resp ct.GetManagerResponse
	var created_at, updated_at, deleted_at sql.NullString
	err := r.db.QueryRow(ctx, `
        SELECT login, fullname, phone, password, salary, branchId, created_at, updated_at, deleted_at
        FROM manager
        WHERE id = $1 and deleted_at IS NULL
    `, req.Id).Scan(
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.BranchId,
		&created_at, &updated_at, &deleted_at,
	)
	resp.Id = req.Id
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting manager by id:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *managerRepo) GetList(ctx context.Context, req *ct.GetListManagerRequest) (*ct.GetListManagerResponse, error) {
	var resp ct.GetListManagerResponse
	var created_at, updated_at, deleted_at sql.NullString

	// Calculate OFFSET based on page and limit
	offset := (req.Page - 1) * req.Limit

	// Initialize filter string
	filter := ""

	// Add search condition if req.Search is not empty
	if req.Search != "" {
		filter = fmt.Sprintf(`
            WHERE (login ILIKE '%%%v%%'
            OR fullname ILIKE '%%%v%%'
            OR phone ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search, req.Search, req.Search)
	} else {
		filter = " WHERE deleted_at IS NULL "
	}

	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
        SELECT id, login, fullname, phone, password, salary, branchId, created_at, updated_at, deleted_at
        FROM manager
        %s
        LIMIT $1 OFFSET $2
    `, filter)

	// Execute the query
	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting manager list:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate response
	for rows.Next() {
		var manager ct.GetManagerResponse
		err := rows.Scan(&manager.Id, &manager.Login, &manager.Fullname, &manager.Phone, &manager.Password, &manager.Salary, &manager.BranchId,
			&created_at, &updated_at, &deleted_at)
		if err != nil {
			log.Println("error while scanning manager row:", err)
			return nil, err
		}
		manager.CreatedAt = created_at.String
		manager.UpdatedAt = updated_at.String
		manager.DeletedAt = deleted_at.String
		resp.Managers = append(resp.Managers, &manager)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over manager rows:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *managerRepo) Update(ctx context.Context, req *ct.UpdateManagerRequest) (*ct.GetManagerResponse, error) {
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	// Perform the update operation
	_, err := r.db.Exec(ctx, `
        UPDATE manager
        SET login = $1, fullname = $2, phone = $3, password = $4, salary = $5, branchId = $6, updated_at = $7
        WHERE id = $8
    `, req.Login, req.Fullname, req.Phone, req.Password, req.Salary, req.BranchId, updated_at, req.Id)

	if err != nil {
		log.Println("error while updating manager:", err)
		return nil, err
	}

	manager, err := r.GetByID(ctx, &ct.ManagerID{Id: req.Id})
	if err != nil {
		log.Println("error while getting manager by id after update")
		return nil, err
	}

	return manager, nil
}

func (r *managerRepo) Delete(ctx context.Context, req *ct.ManagerID) (*ct.ManagerEmpty, error) {
	resp := &ct.ManagerEmpty{}
	deleted_at := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(ctx, `
    UPDATE manager
    SET deleted_at = $1
    WHERE id = $2
`, deleted_at, req.Id)

	if err != nil {
		log.Println("error while deleting manager:", err)
		return nil, err
	}

	resp.Msg = "Successful"
	return resp, nil
}


func (r *managerRepo) GetLastLogin(ctx context.Context) (string, error) {
    var login string
    err := r.db.QueryRow(ctx, `
        SELECT login FROM manager
        ORDER BY login DESC LIMIT 1
    `).Scan(&login)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) || err.Error()=="no rows in result set" {
            log.Println("No rows found, returning default login: M00000")
            return "M00000", nil
        }
        // Log the unexpected error
        log.Printf("Unexpected database error: %v", err)
        return "", fmt.Errorf("database query error: %w", err)
    }

    log.Printf("Retrieved login: %s", login)
    return login, nil
}

func GenerateNewLoginManager(log string) string {
	prefix := "M"
	numbStr := log[1:]
	num, _ := strconv.Atoi(numbStr)
	newNum := num + 1
	return fmt.Sprintf("%s%05d", prefix, newNum)
}


func (r *managerRepo) GetByLogin(ctx context.Context, login string) (*ct.GetManagerResponse, error) {
	var resp ct.GetManagerResponse
	var created_at, updated_at, deleted_at sql.NullString
	err := r.db.QueryRow(ctx, `
        SELECT id,login, fullname, phone, password, salary, branchId, created_at, updated_at, deleted_at
        FROM manager
        WHERE login = $1 and deleted_at IS NULL
    `, login).Scan(&resp.Id,
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.BranchId,
		&created_at, &updated_at, &deleted_at,
	)
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting manager by id:", err)
		return nil, err
	}
	return &resp, nil
}