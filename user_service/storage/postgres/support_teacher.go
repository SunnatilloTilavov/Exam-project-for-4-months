package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	ct "user_service/genproto/user_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type supportteacherRepo struct {
	db *pgxpool.Pool // PostgreSQL Pool
}

func NewSupportteacherRepo(db *pgxpool.Pool) *supportteacherRepo {
	return &supportteacherRepo{
		db: db,
	}
}

// Create teacher
func (r *supportteacherRepo) Create(ctx context.Context, req *ct.CreateSupportTeacherRequest) (*ct.SupportTeacherResponse, error) {
	var id uuid.UUID
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	loginlast, err := r.GetLastLogin(ctx)
	if err != nil {
		log.Println("error while creating Sp teacher:", err)
		return nil, err
	}
	login := GenerateNewLoginSPTeacher(loginlast)
fmt.Println(login)
	err = r.db.QueryRow(ctx, `
		INSERT INTO support_teacher (login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, branchId, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`, login, req.Fullname, req.Phone, req.Password, req.Salary, req.IeltsScore, req.IeltsAttemptCount, req.BranchId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating teacher:", err)
		return nil, err
	}

	return &ct.SupportTeacherResponse{
		Id:                id.String(),
		Login:             login,
		Fullname:          req.Fullname,
		Phone:             req.Phone,
		Password:          req.Password,
		Salary:            req.Salary,
		IeltsScore:        req.IeltsScore,
		IeltsAttemptCount: req.IeltsAttemptCount,
		BranchId:          req.BranchId,
	}, nil
}

// Get teacher by ID
func (r *supportteacherRepo) GetByID(ctx context.Context, req *ct.SupportTeacherID) (*ct.GetSupportTeacherResponse, error) {
	var resp ct.GetSupportTeacherResponse
	var updatedAt, branchId, created_at sql.NullString

	err := r.db.QueryRow(ctx, `
		SELECT login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, branchId, created_at, updated_at
		FROM support_teacher 
		WHERE id = $1 AND  deleted_at IS  NULL
	`, req.Id).Scan(
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.IeltsScore, &resp.IeltsAttemptCount,
		&branchId, &created_at, &updatedAt,
	)
	if err != nil {
		log.Println("error while getting teacher by id:", err)
		return nil, err
	}

	resp.Id = req.Id
	resp.CreatedAt = created_at.String
	resp.BranchId = branchId.String
	resp.UpdatedAt = updatedAt.String
	return &resp, nil
}

// Get list of teachers with pagination and search
func (r *supportteacherRepo) GetList(ctx context.Context, req *ct.GetListSupportTeacherRequest) (*ct.GetListSupportTeacherResponse, error) {
	var resp ct.GetListSupportTeacherResponse
	var updatedAt, branchId, created_at sql.NullString

	// Calculate OFFSET based on page and limit
	offset := (req.Page - 1) * req.Limit

	// Initialize filter string
	filter := ""

	// Add search condition if req.Search is not empty
	if req.Search != "" {
		filter = fmt.Sprintf(`
			WHERE (login ILIKE '%%%v%%'
			OR fullname ILIKE '%%%v%%'
			OR phone ILIKE '%%%v%%') AND deleted_at IS  NULL
		`, req.Search, req.Search, req.Search)
	} else {
		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL`)
	}

	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
		SELECT id, login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, branchId, created_at, updated_at
		FROM support_teacher
		%s
		LIMIT $1 OFFSET $2
	`, filter)

	// Execute the query
	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting teacher list:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate response
	for rows.Next() {
		var teacher ct.GetSupportTeacherResponse
		err := rows.Scan(&teacher.Id, &teacher.Login, &teacher.Fullname, &teacher.Phone, &teacher.Password, &teacher.Salary, &teacher.IeltsScore, &teacher.IeltsAttemptCount, &branchId, &created_at, &updatedAt)
		if err != nil {
			log.Println("error while scanning teacher row:", err)
			return nil, err
		}
		teacher.BranchId = branchId.String
		teacher.CreatedAt = created_at.String
		teacher.UpdatedAt = updatedAt.String
		resp.GetSupportTeachers = append(resp.GetSupportTeachers, &teacher)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over teacher rows:", err)
		return nil, err
	}
	return &resp, nil
}

// Update teacher
func (r *supportteacherRepo) Update(ctx context.Context, req *ct.UpdateSupportTeacherRequest) (*ct.GetSupportTeacherResponse, error) {
	updatedAt := time.Now().Format("2006-01-02 15:04:05")

	// Perform the update operation
	_, err := r.db.Exec(ctx, `
		UPDATE support_teacher
		SET login = $1, fullname = $2, phone = $3, password = $4, salary = $5, 
			ieltsScore = $6, ieltsAttemptCount = $7, branchId = $8, updated_at = $9
		WHERE id = $10
	`, req.Login, req.Fullname, req.Phone, req.Password, req.Salary, req.IeltsScore, req.IeltsAttemptCount, req.BranchId, updatedAt, req.Id)

	if err != nil {
		log.Println("error while updating teacher:", err)
		return nil, err
	}

	// Retrieve the updated teacher record
	teacher, err := r.GetByID(ctx, &ct.SupportTeacherID{Id: req.Id})
	if err != nil {
		log.Println("error while getting teacher by id after update:", err)
		return nil, err
	}

	return teacher, nil
}

// Delete teacher
func (r *supportteacherRepo) Delete(ctx context.Context, req *ct.SupportTeacherID) (*ct.SupportTeacherEmpty, error) {
	deletedAt := time.Now().Format("2006-01-02 15:04:05")
	resp := &ct.SupportTeacherEmpty{}
	// Perform the soft delete operation
	_, err := r.db.Exec(ctx, `
		UPDATE support_teacher
		SET  deleted_at = $1
		WHERE id = $2
	`, deletedAt, req.Id)

	if err != nil {
		log.Println("error while deleting teacher:", err)
		return nil, err
	}
	resp.Msg = "Successful"
	return resp, nil
}

func (r *supportteacherRepo) GetLastLogin(ctx context.Context) (string, error) {
	var login string
	err := r.db.QueryRow(ctx, `
        SELECT login FROM support_teacher
        ORDER BY login DESC LIMIT 1
    `).Scan(&login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "no rows in result set" {
			log.Println("No rows found, returning default login: ST00000")
			return "ST00000", nil
		}
		// Log the unexpected error
		log.Printf("Unexpected database error: %v", err)
		return "", fmt.Errorf("database query error: %w", err)
	}

	log.Printf("Retrieved login: %s", login)
	return login, nil
}

func GenerateNewLoginSPTeacher(log string) string {
	prefix := "ST"
	numbStr := log[2:]
	num, _ := strconv.Atoi(numbStr)
	newNum := num + 1
	return fmt.Sprintf("%s%05d", prefix, newNum)
}


func (r *supportteacherRepo) GetByLogin(ctx context.Context, login string) (*ct.GetSupportTeacherResponse, error) {
	var resp ct.GetSupportTeacherResponse
	var updatedAt, branchId, created_at sql.NullString

	err := r.db.QueryRow(ctx, `
		SELECT id,login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, branchId, created_at, updated_at
		FROM support_teacher 
		WHERE id = $1 AND  deleted_at IS  NULL
	`,login ).Scan(&resp.Id,
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.IeltsScore, &resp.IeltsAttemptCount,
		&branchId, &created_at, &updatedAt,
	)
	if err != nil {
		log.Println("error while getting teacher by id:", err)
		return nil, err
	}
	resp.CreatedAt = created_at.String
	resp.BranchId = branchId.String
	resp.UpdatedAt = updatedAt.String
	return &resp, nil
}



func (r *supportteacherRepo) GetReportList(ctx context.Context, req *ct.GetReportListSupportTeacherRequest) (*ct.GetReportListSupportTeacherResponse, error) {
	var resp ct.GetReportListSupportTeacherResponse
	var updatedAt, branchId, created_at sql.NullString

	// Calculate OFFSET based on page and limit
	offset := (req.Page - 1) * req.Limit

	// Initialize filter string
	filter := ""

	// Add search condition if req.Search is not empty
	if req.Search != "" {
		filter = fmt.Sprintf(`
			WHERE (login ILIKE '%%%v%%'
			OR fullname ILIKE '%%%v%%'
			OR phone ILIKE '%%%v%%') AND deleted_at IS  NULL
		`, req.Search, req.Search, req.Search)
	} else {
		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL`)
	}

	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
		SELECT id, login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, branchId, created_at, updated_at
		FROM support_teacher
		%s
		LIMIT $1 OFFSET $2
	`, filter)

	// Execute the query
	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting teacher list:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate response
	for rows.Next() {
		var teacher ct.GetReportSupportTeacherResponse
		err := rows.Scan(&teacher.Id, &teacher.Login, &teacher.Fullname, &teacher.Phone, &teacher.Password, &teacher.Salary, &teacher.IeltsScore, &teacher.IeltsAttemptCount, &branchId, &created_at, &updatedAt)
		if err != nil {
			log.Println("error while scanning teacher row:", err)
			return nil, err
		}
		teacher.BranchId = branchId.String
		teacher.CreatedAt = created_at.String
		teacher.UpdatedAt = updatedAt.String

// Parse created date to calculate days worked
startDateStr := teacher.CreatedAt[:10]
startDate, err := time.Parse("2006-01-02", startDateStr)
if err != nil {
	log.Println("Error parsing start date:", err)
	return nil, err
}

// Calculate days worked since startDate
daysWorked := time.Since(startDate).Hours() / 24

// Calculate monthly rate based on salary (assuming salary is daily rate)
monthlyRate := float64(teacher.Salary)  // Adjust as per your salary logic

// Calculate total sum based on days worked and monthly rate
totalSum := int64(daysWorked * monthlyRate)

// Assign totalSum as string to Totalsum field
teacher.Totalsum = fmt.Sprintf("%d", totalSum)

// Assign other fields to response struct






		resp.GetSupportTeachers = append(resp.GetSupportTeachers, &teacher)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over teacher rows:", err)
		return nil, err
	}
	return &resp, nil
}
