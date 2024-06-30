package postgres

import (
	"context"
	"database/sql"

	"fmt"
	"log"
	"strconv"
	"time"
	"errors"

	ct "user_service/genproto/user_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type administrationRepo struct {
	db *pgxpool.Pool
}

func NewAdministrationRepo(db *pgxpool.Pool) *administrationRepo {
	return &administrationRepo{
		db: db,
	}
}

func (r *administrationRepo) Create(ctx context.Context, req *ct.CreateAdministrationRequest) (*ct.AdministrationResponse, error) {
	var id uuid.UUID

	loginlast, err := r.GetLastLogin(ctx)
	if err != nil {
		log.Println("error while creating administration:", err)
		return nil, err
	}
	login:=GenerateNewLogin(loginlast)

	createdAt := time.Now().Format("2006-01-02 15:04:05")

	err = r.db.QueryRow(ctx, `
        INSERT INTO administration (login, fullname, phone, password, salary, ieltsScore, branchId, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `, login, req.Fullname, req.Phone, req.Password, req.Salary, req.IeltsScore, req.BranchId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating administration:", err)
		return nil, err
	}

	return &ct.AdministrationResponse{
		Id:         id.String(),
		Login:      login,
		Fullname:   req.Fullname,
		Phone:      req.Phone,
		Password:   req.Password,
		Salary:     req.Salary,
		IeltsScore: req.IeltsScore,
		BranchId:   req.BranchId,
	}, nil
}

func (r *administrationRepo) GetByID(ctx context.Context, req *ct.AdministrationID) (*ct.GetAdministrationResponse, error) {
	var resp ct.GetAdministrationResponse
	var updated_at, branchId, created_at, deleted_at sql.NullString

	err := r.db.QueryRow(ctx, `
        SELECT login, fullname, phone, password, salary, ieltsScore, branchId, created_at, updated_at, deleted_at
        FROM administration 
        WHERE id = $1 AND deleted_at IS  NULL
    `, req.Id).Scan(
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.IeltsScore,
		&branchId, &created_at, &updated_at, &deleted_at,
	)
	resp.UpdatedAt = updated_at.String
	resp.BranchId = branchId.String
	resp.CreatedAt = created_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting administration by id:", err)
		return nil, err
	}
	resp.Id = req.Id
	return &resp, nil
}

func (r *administrationRepo) GetList(ctx context.Context, req *ct.GetListAdministrationRequest) (*ct.GetListAdministrationResponse, error) {
	var resp ct.GetListAdministrationResponse
	var updated_at, branchId, created_at, deleted_at sql.NullString

	offset := (req.Page - 1) * req.Limit

	filter := ""
	if req.Search != "" {
		filter = fmt.Sprintf(`
            WHERE (login ILIKE '%%%v%%'
            OR fullname ILIKE '%%%v%%'
            OR phone ILIKE '%%%v%%') AND deleted_at IS  NULL
        `, req.Search, req.Search, req.Search)
	} else {
		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL `)
	}

	query := fmt.Sprintf(`
        SELECT id, login, fullname, phone, password, salary, ieltsScore, branchId, created_at, updated_at, deleted_at
        FROM administration
        %s
        LIMIT $1 OFFSET $2
    `, filter)

	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting administration list:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var administration ct.GetAdministrationResponse
		err := rows.Scan(&administration.Id, &administration.Login, &administration.Fullname, &administration.Phone, &administration.Password, &administration.Salary, &administration.IeltsScore, &branchId, &created_at, &updated_at, &deleted_at)
		if err != nil {
			log.Println("error while scanning administration row:", err)
			return nil, err
		}
		administration.UpdatedAt = updated_at.String
		administration.BranchId = branchId.String
		administration.CreatedAt = created_at.String
		administration.DeletedAt = deleted_at.String
		resp.Getadministrations = append(resp.Getadministrations, &administration)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over administration rows:", err)
		return nil, err
	}
	resp.Count = int64(len(resp.Getadministrations))
	return &resp, nil
}

func (r *administrationRepo) Update(ctx context.Context, req *ct.UpdateAdministrationRequest) (*ct.GetAdministrationResponse, error) {
	updated_at := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
        UPDATE administration
        SET  fullname = $1, phone = $2, password = $3, salary = $4, ieltsScore = $5, branchId = $6, updated_at = $7
        WHERE id = $8
    `,  req.Fullname, req.Phone, req.Password, req.Salary, req.IeltsScore, req.BranchId, updated_at, req.Id)

	if err != nil {
		log.Println("error while updating administration:", err)
		return nil, err
	}

	administration, err := r.GetByID(ctx, &ct.AdministrationID{Id: req.Id})
	if err != nil {
		log.Println("error while getting administration by id after update:", err)
		return nil, err
	}

	return administration, nil
}

func (r *administrationRepo) Delete(ctx context.Context, req *ct.AdministrationID) (*ct.AdministrationEmpty, error) {
	resp := &ct.AdministrationEmpty{}
	deleted_at := time.Now().Format("2006-01-02 15:04:05")

	_, err := r.db.Exec(ctx, `
        UPDATE administration
        SET deleted_at = $1
        WHERE id = $2
    `, deleted_at, req.Id)

	if err != nil {
		log.Println("error while deleting administration:", err)
		return nil, err
	}

	resp.Msg = "Successful"
	return resp, nil
}


func (r *administrationRepo) GetLastLogin(ctx context.Context) (string, error) {
    var login string
    err := r.db.QueryRow(ctx, `
        SELECT login FROM administration 
        ORDER BY login DESC LIMIT 1
    `).Scan(&login)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) || err.Error()=="no rows in result set" {
            log.Println("No rows found, returning default login: A00000")
            return "A00000", nil
        }
        // Log the unexpected error
        log.Printf("Unexpected database error: %v", err)
        return "", fmt.Errorf("database query error: %w", err)
    }

    log.Printf("Retrieved login: %s", login)
    return login, nil
}

func GenerateNewLogin(log string) string {
	prefix := "A"
	numbStr := log[1:]
	num, _ := strconv.Atoi(numbStr)
	newNum := num + 1
	return fmt.Sprintf("%s%05d", prefix, newNum)
}


func (r *administrationRepo) GetByLogin(ctx context.Context, login string) (*ct.GetAdministrationResponse, error) {
	var resp ct.GetAdministrationResponse
	var updated_at, branchId, created_at, deleted_at sql.NullString

	err := r.db.QueryRow(ctx, `
        SELECT id,login, fullname, phone, password, salary, ieltsScore, branchId, created_at, updated_at, deleted_at
        FROM administration 
        WHERE login = $1 AND deleted_at IS  NULL
    `, login).Scan(&resp.Id,
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.IeltsScore,
		&branchId, &created_at, &updated_at, &deleted_at,
	)
	resp.UpdatedAt = updated_at.String
	resp.BranchId = branchId.String
	resp.CreatedAt = created_at.String
	resp.DeletedAt = deleted_at.String
	
	if err != nil {
		log.Println("error while getting administration by id:", err)
		return nil, err
	}
	return &resp, nil
}




// func (r *administrationRepo) GetReportList(ctx context.Context, req *ct.GetReportListAdministrationRequest) (*ct.GetReportListAdministrationResponse, error) {
// 	var resp ct.GetReportListAdministrationResponse
// 	var updated_at, branchId, created_at, deleted_at sql.NullString

// 	offset := (req.Page - 1) * req.Limit

// 	filter := ""
// 	if req.Search != "" {
// 		filter = fmt.Sprintf(`
//             WHERE (login ILIKE '%%%v%%'
//             OR fullname ILIKE '%%%v%%'
//             OR phone ILIKE '%%%v%%') AND deleted_at IS  NULL
//         `, req.Search, req.Search, req.Search)
// 	} else {
// 		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL `)
// 	}

// 	query := fmt.Sprintf(`
//         SELECT id, login, fullname, phone, password, salary, ieltsScore, branchId, created_at, updated_at, deleted_at
//         FROM administration
//         %s
//         LIMIT $1 OFFSET $2
//     `, filter)

// 	rows, err := r.db.Query(ctx, query, req.Limit, offset)
// 	if err != nil {
// 		log.Println("error while getting administration list:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var administration ct.GetReportAdministrationResponse
// 		err := rows.Scan(&administration.Id, &administration.Login, &administration.Fullname, &administration.Phone, &administration.Password, &administration.Salary, &administration.IeltsScore, &branchId, &created_at, &updated_at, &deleted_at)
// 		if err != nil {
// 			log.Println("error while scanning administration row:", err)
// 			return nil, err
// 		}

// 		administration.UpdatedAt = updated_at.String
// 		administration.BranchId = branchId.String
// 		administration.CreatedAt = created_at.String
// 		administration.DeletedAt = deleted_at.String


// 		startDateStr := administration.CreatedAt[:10]
// 		startDate, err := time.Parse("2006-01-02", startDateStr)
// 		if err != nil {
// 			fmt.Println("Error parsing start date:", err)
// 			return nil, err
// 		}
	
// 		daysWorked := int64(time.Since(startDate).Hours() / 24)

// 		monthlyRate:= administration.Salary
	
// 		administration.Totalsum = string(daysWorked * monthlyRate)


// 		resp.Getadministrations = append(resp.Getadministrations, &administration)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Println("error after iterating over administration rows:", err)
// 		return nil, err
// 	}
// 	resp.Count = int64(len(resp.Getadministrations))
// 	return &resp, nil
// }

func (r *administrationRepo) GetReportList(ctx context.Context, req *ct.GetReportListAdministrationRequest) (*ct.GetReportListAdministrationResponse, error) {
	var resp ct.GetReportListAdministrationResponse
	var updated_at, branchId, created_at, deleted_at sql.NullString

	offset := (req.Page - 1) * req.Limit

	filter := ""
	if req.Search != "" {
		filter = fmt.Sprintf(`
            WHERE (login ILIKE '%%%v%%'
            OR fullname ILIKE '%%%v%%'
            OR phone ILIKE '%%%v%%') AND deleted_at IS  NULL
        `, req.Search, req.Search, req.Search)
	} else {
		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL `)
	}

	query := fmt.Sprintf(`
        SELECT id, login, fullname, phone, password, salary, ieltsScore, branchId, created_at, updated_at, deleted_at
        FROM administration
        %s
        LIMIT $1 OFFSET $2
    `, filter)

	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting administration list:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var administration ct.GetReportAdministrationResponse
		err := rows.Scan(&administration.Id, &administration.Login, &administration.Fullname, &administration.Phone, &administration.Password, &administration.Salary, &administration.IeltsScore, &branchId, &created_at, &updated_at, &deleted_at)
		if err != nil {
			log.Println("error while scanning administration row:", err)
			return nil, err
		}

		administration.UpdatedAt = updated_at.String
		administration.BranchId = branchId.String
		administration.CreatedAt = created_at.String
		administration.DeletedAt = deleted_at.String

		// Parse created date to calculate days worked
		startDateStr := administration.CreatedAt[:10]
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			log.Println("Error parsing start date:", err)
			return nil, err
		}

		daysWorked := time.Since(startDate).Hours() / 24

		monthlyRate := float64(administration.Salary/30)
		totalSum := int64(daysWorked * monthlyRate)

	
		administration.Totalsum = fmt.Sprintf("%d", totalSum)

		resp.Getadministrations = append(resp.Getadministrations, &administration)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over administration rows:", err)
		return nil, err
	}
	resp.Count = int64(len(resp.Getadministrations))
	return &resp, nil
}
