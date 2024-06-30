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

type studentRepo struct {
	db *pgxpool.Pool // Assuming you are using pgx/v4
}

func NewStudentRepo(db *pgxpool.Pool) *studentRepo {
	return &studentRepo{
		db: db,
	}
}

func (r *studentRepo) Create(ctx context.Context, req *ct.CreateStudentRequest) (*ct.StudentResponse, error) {
	var id uuid.UUID

	loginlast, err := r.GetLastLogin(ctx)
	if err != nil {
		log.Println("error while creating student:", err)
		return nil, err
	}
	login:=GenerateNewLoginStudent(loginlast)


	createdAt := time.Now().Format("2006-01-02 15:04:05")

	err = r.db.QueryRow(ctx, `
        INSERT INTO student (login, fullname, phone, password, groupName, branchId, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id
    `, login, req.Fullname, req.Phone, req.Password, req.GroupName, req.BranchId, createdAt).Scan(&id)
	if err != nil {
		log.Println("error while creating student:", err)
		return nil, err
	}

	return &ct.StudentResponse{
		Id:        id.String(),
		Login:     login,
		Fullname:  req.Fullname,
		Phone:     req.Phone,
		Password:  req.Password,
		GroupName: req.GroupName,
		BranchId:  req.BranchId,
		CreatedAt: createdAt,
		UpdatedAt: "",
		DeletedAt: "",
	}, nil
}

func (r *studentRepo) GetByID(ctx context.Context, req *ct.StudentID) (*ct.GetStudentResponse, error) {
	var resp ct.GetStudentResponse
	var created_at, updated_at, deleted_at sql.NullString
	err := r.db.QueryRow(ctx, `
        SELECT login, fullname, phone, password, groupName, branchId, created_at, updated_at, deleted_at
        FROM student
        WHERE id = $1
    `, req.Id).Scan(
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.GroupName, &resp.BranchId,
		&created_at, &updated_at, &deleted_at,
	)
	resp.Id = req.Id
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *studentRepo) GetList(ctx context.Context, req *ct.GetListStudentRequest) (*ct.GetListStudentResponse, error) {
	var resp ct.GetListStudentResponse
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
            OR phone ILIKE '%%%v%%'
            OR groupName ILIKE '%%%v%%') AND deleted_at IS NULL
        `, req.Search, req.Search, req.Search, req.Search)
	} else {
		filter = " WHERE deleted_at IS NULL "
	}

	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
        SELECT id, login, fullname, phone, password, groupName, branchId, created_at, updated_at, deleted_at
        FROM student
        %s
        LIMIT $1 OFFSET $2
    `, filter)

	// Execute the query
	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting student list:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate response
	for rows.Next() {
		var student ct.GetStudentResponse
		err := rows.Scan(&student.Id, &student.Login, &student.Fullname, &student.Phone, &student.Password, &student.GroupName,
			&student.BranchId, &created_at, &updated_at, &deleted_at)
		if err != nil {
			log.Println("error while scanning student row:", err)
			return nil, err
		}
		student.CreatedAt = created_at.String
		student.UpdatedAt = updated_at.String
		student.DeletedAt = deleted_at.String
		resp.Students = append(resp.Students, &student)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over student rows:", err)
		return nil, err
	}
	return &resp, nil
}

func (r *studentRepo) Update(ctx context.Context, req *ct.UpdateStudentRequest) (*ct.GetStudentResponse, error) {
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	// Perform the update operation
	_, err := r.db.Exec(ctx, `
        UPDATE student
        SET login = $1, fullname = $2, phone = $3, password = $4, groupName = $5, branchId = $6, updated_at = $7
        WHERE id = $8
    `, req.Login, req.Fullname, req.Phone, req.Password, req.GroupName, req.BranchId, updated_at, req.Id)

	if err != nil {
		log.Println("error while updating student:", err)
		return nil, err
	}

	student, err := r.GetByID(ctx, &ct.StudentID{Id: req.Id})
	if err != nil {
		log.Println("error while getting student by id after update")
		return nil, err
	}

	return student, nil
}

func (r *studentRepo) Delete(ctx context.Context, req *ct.StudentID) (*ct.StudentEmpty, error) {
	resp := &ct.StudentEmpty{}
	deleted_at := time.Now().Format("2006-01-02 15:04:05")
	_, err := r.db.Exec(ctx, `
    UPDATE student
    SET deleted_at = $1
    WHERE id = $2
`, deleted_at, req.Id)

	if err != nil {
		log.Println("error while deleting student:", err)
		return nil, err
	}

	resp.Msg = "Successful"
	return resp, nil
}

func (r *studentRepo) GetLastLogin(ctx context.Context) (string, error) {
	var login string
	err := r.db.QueryRow(ctx, `
        SELECT login FROM student
        ORDER BY login DESC LIMIT 1
    `).Scan(&login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "no rows in result set" {
			log.Println("No rows found, returning default login: S00000")
			return "S00000", nil
		}
		// Log the unexpected error
		log.Printf("Unexpected database error: %v", err)
		return "", fmt.Errorf("database query error: %w", err)
	}

	log.Printf("Retrieved login: %s", login)
	return login, nil
}

func GenerateNewLoginStudent(log string) string {
	prefix := "S"
	numbStr := log[1:]
	num, _ := strconv.Atoi(numbStr)
	newNum := num + 1
	return fmt.Sprintf("%s%05d", prefix, newNum)
}




func (r *studentRepo) GetByLogin(ctx context.Context, login string) (*ct.GetStudentResponse, error) {
	var resp ct.GetStudentResponse
	var created_at, updated_at, deleted_at sql.NullString
	err := r.db.QueryRow(ctx, `
        SELECT id,login, fullname, phone, password, groupName, branchId, created_at, updated_at, deleted_at
        FROM student
        WHERE login = $1
    `, login).Scan(&resp.Id,
		&resp.Login, &resp.Fullname, &resp.Phone, &resp.Password, &resp.GroupName, &resp.BranchId,
		&created_at, &updated_at, &deleted_at,
	)
	resp.CreatedAt = created_at.String
	resp.UpdatedAt = updated_at.String
	resp.DeletedAt = deleted_at.String
	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}
	return &resp, nil
}



// func (r *studentRepo) GetReportList(ctx context.Context, req *ct.GetReportListStudentRequest) (*ct.GetReportListStudentResponse, error) {
// 	var resp ct.GetReportListStudentResponse
// 	var created_at, updated_at, deleted_at sql.NullString

// 	// Calculate OFFSET based on page and limit
// 	offset := (req.Page - 1) * req.Limit

// 	// Initialize filter string
// 	filter := ""

// 	// Add search condition if req.Search is not empty
// 	if req.Search != "" {
// 		filter = fmt.Sprintf(`
//             WHERE (login ILIKE '%%%v%%'
//             OR fullname ILIKE '%%%v%%'
//             OR phone ILIKE '%%%v%%'
//             OR groupName ILIKE '%%%v%%') AND deleted_at IS NULL
//         `, req.Search, req.Search, req.Search, req.Search)
// 	} else {
// 		filter = " WHERE deleted_at IS NULL "
// 	}

// 	// Construct the final query with filter, limit, and offset
// 	query := fmt.Sprintf(`
//         SELECT 
//     s.id,
//     s.login,
//     s.fullname,
//     s.phone,
//     s.password,
//     s.groupName,
//     s.branchId,
//     s.created_at,
//     s.updated_at,
//     s.deleted_at,
//     COALESCE(SUM(sp.paidSum), 0) AS total_paid_sum
// FROM 
//     student s
// LEFT JOIN 
//     student_payment sp ON s.id = sp.studentId
// GROUP BY 
//     s.id,
//     s.login,
//     s.fullname,
//     s.phone,
//     s.password,
//     s.groupName,
//     s.branchId,
//     s.created_at,
//     s.updated_at,
//     s.deleted_at;

//         %s
//         LIMIT $1 OFFSET $2
//     `, filter)

// 	// Execute the query
// 	rows, err := r.db.Query(ctx, query, req.Limit, offset)
// 	if err != nil {
// 		log.Println("error while getting student list:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	// Iterate over the rows and populate response
// 	for rows.Next() {
// 		var student ct.GetReportStudentResponse
// 		err := rows.Scan(&student.Id, &student.Login, &student.Fullname, &student.Phone, &student.Password, &student.GroupName,
// 			&student.BranchId, &created_at, &updated_at, &deleted_at,&student.Paidsum)
// 		if err != nil {
// 			log.Println("error while scanning student row:", err)
// 			return nil, err
// 		}
// 		student.CreatedAt = created_at.String
// 		student.UpdatedAt = updated_at.String
// 		student.DeletedAt = deleted_at.String
// 		resp.Students = append(resp.Students, &student)
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Println("error after iterating over student rows:", err)
// 		return nil, err
// 	}
// 	return &resp, nil
// }

func (r *studentRepo) GetReportList(ctx context.Context, req *ct.GetReportListStudentRequest) (*ct.GetReportListStudentResponse, error) {
	var resp ct.GetReportListStudentResponse
	var created_at, updated_at, deleted_at sql.NullString

	// Calculate OFFSET based on page and limit
	offset := (req.Page - 1) * req.Limit

	// Initialize filter string
	filter := ""

	// Add search condition if req.Search is not empty
	if req.Search != "" {
		filter = fmt.Sprintf(`
            WHERE (s.login ILIKE '%%%v%%'
            OR s.fullname ILIKE '%%%v%%'
            OR s.phone ILIKE '%%%v%%'
            OR s.groupName ILIKE '%%%v%%') AND s.deleted_at IS NULL
        `, req.Search, req.Search, req.Search, req.Search)
	} else {
		filter = " WHERE s.deleted_at IS NULL "
	}

	// Construct the final query with filter, limit, and offset
	query := fmt.Sprintf(`
        SELECT 
            s.id,
            s.login,
            s.fullname,
            s.phone,
            s.password,
            s.groupName,
            s.branchId,
            s.created_at,
            s.updated_at,
            s.deleted_at,
            COALESCE(SUM(sp.paidSum), 0) AS total_paid_sum
        FROM 
            student s
        LEFT JOIN 
            student_payment sp ON s.id = sp.studentId
        %s
        GROUP BY 
            s.id,
            s.login,
            s.fullname,
            s.phone,
            s.password,
            s.groupName,
            s.branchId,
            s.created_at,
            s.updated_at,
            s.deleted_at
        LIMIT $1 OFFSET $2
    `, filter)

	// Execute the query
	rows, err := r.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error while getting student list:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate response
	for rows.Next() {
		var student ct.GetReportStudentResponse
		err := rows.Scan(&student.Id, &student.Login, &student.Fullname, &student.Phone, &student.Password, &student.GroupName,
			&student.BranchId, &created_at, &updated_at, &deleted_at, &student.Paidsum)
		if err != nil {
			log.Println("error while scanning student row:", err)
			return nil, err
		}
		student.CreatedAt = created_at.String
		student.UpdatedAt = updated_at.String
		student.DeletedAt = deleted_at.String
		resp.Students = append(resp.Students, &student)
	}
	if err := rows.Err(); err != nil {
		log.Println("error after iterating over student rows:", err)
		return nil, err
	}
	return &resp, nil
}
