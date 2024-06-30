package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"strconv"
	"errors"

	ct "user_service/genproto/user_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewTeacherRepo(db *pgxpool.Pool) *teacherRepo {
	return &teacherRepo{
		db: db,
	}
}

func (r *teacherRepo) Create(ctx context.Context, req *ct.CreateTeacherRequest) (*ct.TeacherResponse, error) {
	var id uuid.UUID
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	loginlast, err := r.GetLastLogin(ctx)
	if err != nil {
		log.Println("error while creating teacher:", err)
		return nil, err
	}
	login := GenerateNewLoginTeacher(loginlast)


	err = r.db.QueryRow(ctx, `
        INSERT INTO teacher (login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, supportTeacherId, branchId, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING id
    `, login, req.Fullname, req.Phone, req.Password, req.Salary, req.IeltsScore, req.IeltsAttemptCount, req.SupportTeacherId, req.BranchId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating teacher:", err)
		return nil, err
	}

	return &ct.TeacherResponse{
		Id:                id.String(),
		Login:             login,
		Fullname:          req.Fullname,
		Phone:             req.Phone,
		Password:          req.Password,
		Salary:            req.Salary,
		IeltsScore:        req.IeltsScore,
		IeltsAttemptCount: req.IeltsAttemptCount,
		SupportTeacherId:  req.SupportTeacherId,
		BranchId:          req.BranchId,
	}, nil
}

func (r *teacherRepo) GetByID(ctx context.Context, req *ct.TeacherID) (*ct.GetTeacherResponse, error) {
	var resp ct.GetTeacherResponse
	var updated_at, branchId, created_at sql.NullString
	err := r.db.QueryRow(ctx, `
        SELECT login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, supportTeacherId, branchId, created_at, updated_at
        FROM teacher
        WHERE id = $1 AND deleted_at IS  NULL
    `, req.Id).Scan(
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.IeltsScore, &resp.IeltsAttemptCount, &resp.SupportTeacherId,
		&branchId, &created_at, &updated_at,
	)
	resp.UpdatedAt = updated_at.String
	resp.BranchId = branchId.String
	resp.CreatedAt = created_at.String
	if err != nil {
		log.Println("error while getting teacher by id:", err)
		return nil, err
	}
	resp.Id = req.Id
	return &resp, nil
}
func (r *teacherRepo) GetList(ctx context.Context, req *ct.GetListTeacherRequest) (*ct.GetListTeacherResponse, error) {
	var resp ct.GetListTeacherResponse
	var updated_at, branchId, created_at sql.NullString

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
	}else {
		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL `)
	}


	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
        SELECT id, login, fullname, phone, password, salary,
         ieltsScore, ieltsAttemptCount, 
         supportTeacherId, branchId, created_at, updated_at
        FROM teacher
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
		var teacher ct.GetTeacherResponse
		err := rows.Scan(&teacher.Id, &teacher.Login, &teacher.Fullname, &teacher.Phone, &teacher.Password, &teacher.Salary, &teacher.IeltsScore, &teacher.IeltsAttemptCount, &teacher.SupportTeacherId, &branchId, &created_at, &updated_at)
		if err != nil {
			log.Println("error while scanning teacher row:", err)
			return nil, err
		}
		teacher.UpdatedAt = updated_at.String
		teacher.BranchId = branchId.String
		teacher.CreatedAt = created_at.String
		resp.GetTeacherResponse = append(resp.GetTeacherResponse, &teacher)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over teacher rows:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *teacherRepo) Update(ctx context.Context, req *ct.UpdateTeacherRequest) (*ct.GetTeacherResponse, error) {
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	// Perform the update operation
	_, err := r.db.Exec(ctx, `
        UPDATE teacher
        SET login = $1, fullname = $2, phone = $3, password = $4, salary = $5, 
            ieltsScore = $6, ieltsAttemptCount = $7, supportTeacherId = $8, 
            branchId = $9, updated_at = $10
        WHERE id = $11
    `, req.Login, req.Fullname, req.Phone, req.Password, req.Salary, req.IeltsScore,
		req.IeltsAttemptCount, req.SupportTeacherId, req.BranchId, updated_at, req.Id)

	if err != nil {
		log.Println("error while updating teacher:", err)
		return nil, err
	}

	teacher, err := r.GetByID(ctx, &ct.TeacherID{Id: req.Id})
	if err != nil {
		log.Println("error while getting category by id after update")
		return nil, err
	}

	return teacher, nil

}

func (r *teacherRepo) Delete(ctx context.Context, req *ct.TeacherID) (*ct.TeacherEmpty, error) {
	resp:=&ct.TeacherEmpty{}
    deleted_at := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(ctx, `
    UPDATE teacher
    SET updated_at = $1
    WHERE id = $2
`, deleted_at, req.Id)

	if err != nil {
		log.Println("error while updating teacher:", err)
		return nil, err
	}
    
    resp.Msg="Successful"

	return resp, nil
}


func (r *teacherRepo) GetLastLogin(ctx context.Context) (string, error) {
	var login string
	err := r.db.QueryRow(ctx, `
        SELECT login FROM teacher 
        ORDER BY login DESC LIMIT 1
    `).Scan(&login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "no rows in result set" {
			log.Println("No rows found, returning default login: T00000")
			return "T00000", nil
		}
		// Log the unexpected error
		log.Printf("Unexpected database error: %v", err)
		return "", fmt.Errorf("database query error: %w", err)
	}

	log.Printf("Retrieved login: %s", login)
	return login, nil
}

func GenerateNewLoginTeacher(log string) string {
	prefix := "T"
	numbStr := log[1:]
	num, _ := strconv.Atoi(numbStr)
	newNum := num + 1
	return fmt.Sprintf("%s%05d", prefix, newNum)
}

func (r *teacherRepo) GetByLogin(ctx context.Context, login string) (*ct.GetTeacherResponse, error) {
	var resp ct.GetTeacherResponse
	var updated_at, branchId, created_at sql.NullString
	err := r.db.QueryRow(ctx, `
        SELECT id,login, fullname, phone, password, salary, ieltsScore, ieltsAttemptCount, supportTeacherId, branchId, created_at, updated_at
        FROM teacher
        WHERE login = $1 AND deleted_at IS  NULL
    `, login).Scan(&resp.Id,
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary, &resp.IeltsScore, &resp.IeltsAttemptCount, &resp.SupportTeacherId,
		&branchId, &created_at, &updated_at,
	)
	resp.UpdatedAt = updated_at.String
	resp.BranchId = branchId.String
	resp.CreatedAt = created_at.String
	if err != nil {
		log.Println("error while getting teacher by id:", err)
		return nil, err
	}
	return &resp, nil
}


// func (r *teacherRepo) GetReportList(ctx context.Context, req *ct.GetReportListTeacherRequest) (*ct.GetReportListTeacherResponse, error) {
// 	var resp ct.GetReportListTeacherResponse
// 	var updated_at, branchId, created_at sql.NullString

// 	// Calculate OFFSET based on page and limit
// 	offset := (req.Page - 1) * req.Limit

// 	// Initialize filter string
// 	filter := ""

// 	// Add search condition if req.Search is not empty
// 	if req.Search != "" {
// 		filter = fmt.Sprintf(`
//             WHERE (login ILIKE '%%%v%%'
//             OR fullname ILIKE '%%%v%%'
//             OR phone ILIKE '%%%v%%') AND deleted_at IS  NULL
//         `, req.Search, req.Search, req.Search)
// 	}else {
// 		filter = fmt.Sprintf(` WHERE  deleted_at IS  NULL `)
// 	}


// 	// Construct the final query with filter, limit, and offset
// 	query := fmt.Sprintf(`
//         SELECT id, login, fullname, phone, password, salary,
//          ieltsScore, ieltsAttemptCount, 
//          supportTeacherId, branchId, created_at, updated_at
//         FROM teacher
//         %s
//         LIMIT $1 OFFSET $2
//     `, filter)

// 	// Execute the query
// 	rows, err := r.db.Query(ctx, query, req.Limit, offset)
// 	if err != nil {
// 		log.Println("error while getting teacher list:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Iterate over the rows and populate response
// 	for rows.Next() {
// 		var teacher ct.GetReportTeacherResponse
// 		err := rows.Scan(&teacher.Id, &teacher.Login, &teacher.Fullname, &teacher.Phone, &teacher.Password, &teacher.Salary, &teacher.IeltsScore, &teacher.IeltsAttemptCount, &teacher.SupportTeacherId, &branchId, &created_at, &updated_at)
// 		if err != nil {
// 			log.Println("error while scanning teacher row:", err)
// 			return nil, err
// 		}
// 		startDateStr := teacher.CreatedAt[:10]
// 		startDate, err := time.Parse("2006-01-02", startDateStr)
// 		if err != nil {
// 			log.Println("Error parsing start date:", err)
// 			return nil, err
// 		}

// 		daysWorked := time.Since(startDate).Hours() / 24

// 		monthlyRate := float64(teacher.Salary/30)
// 		totalSum := int64(daysWorked * monthlyRate)

	
// 		teacher.Totalsum = fmt.Sprintf("%d", totalSum)



// 		teacher.UpdatedAt = updated_at.String
// 		teacher.BranchId = branchId.String
// 		teacher.CreatedAt = created_at.String
// 		resp.GetTeacherResponse = append(resp.GetTeacherResponse, &teacher)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Println("error after iterating over teacher rows:", err)
// 		return nil, err
// 	}
// 	return &resp, nil
// }
func (r *teacherRepo) GetReportList(ctx context.Context, req *ct.GetReportListTeacherRequest) (*ct.GetReportListTeacherResponse, error) {
	var resp ct.GetReportListTeacherResponse
	var updated_at, branchId, created_at sql.NullString

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
		filter = fmt.Sprintf(` WHERE deleted_at IS NULL `)
	}

	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
        SELECT id, login, fullname, phone, password, salary,
         ieltsScore, ieltsAttemptCount, supportTeacherId, branchId, created_at, updated_at
        FROM teacher
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
		var teacher ct.GetReportTeacherResponse
		err := rows.Scan(&teacher.Id, &teacher.Login, &teacher.Fullname, &teacher.Phone, &teacher.Password, &teacher.Salary, &teacher.IeltsScore, &teacher.IeltsAttemptCount, &teacher.SupportTeacherId, &branchId, &created_at, &updated_at)
		if err != nil {
			log.Println("error while scanning teacher row:", err)
			return nil, err
		}

		teacher.UpdatedAt = updated_at.String
		teacher.BranchId = branchId.String
		teacher.CreatedAt = created_at.String

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
	

		// Append teacher response to the response list
		resp.GetTeacherResponse = append(resp.GetTeacherResponse, &teacher)
	}

	// Check for any error after iterating over rows
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over teacher rows:", err)
		return nil, err
	}

	// Return the populated response
	return &resp, nil
}
