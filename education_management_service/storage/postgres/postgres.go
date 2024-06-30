package postgres

import (
	"context"
	"education_management_service/config"
	"education_management_service/storage"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db     *pgxpool.Pool
	branch storage.BranchRepoI
	group  storage.GroupRepoI
	jurnal storage.JurnalRepoI
	schedule storage.ScheduleRepoI
	task storage.TaskRepoI
	studenttask  storage.StudentTaskRepoI
	event storage.EventRepoI
	eventStudent  storage.EventStudentRepoI
	studentpayment storage.StudentPaymentRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Branch() storage.BranchRepoI {
	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}
	return s.branch
}

func (s *Store) Group() storage.GroupRepoI {
	if s.group == nil {
		s.group = NewGroupRepo(s.db)
	}
	return s.group
}

func (s *Store) Jurnal() storage.JurnalRepoI {
	if s.jurnal == nil {
		s.jurnal = NewJurnalRepo(s.db)
	}
	return s.jurnal
}

func (s *Store) Schedule() storage.ScheduleRepoI {
	if s.schedule == nil {
		s.schedule = NewScheduleRepo(s.db)
	}
	return s.schedule
}


// Add this method to return the db pool
func (s *Store) DB() *pgxpool.Pool {
	return s.db
}

func (s *Store) Task() storage.TaskRepoI {
	if s.task == nil {
		s.task = NewTaskRepo(s.db)
	}
	return s.task
}

func (s *Store) StudentTask() storage.StudentTaskRepoI {
	if s.studenttask == nil {
		s.studenttask = NewStudentTaskRepo(s.db)
	}
	return s.studenttask
}

func (s *Store) Event() storage.EventRepoI {
	if s.event == nil {
		s.event = NewEventRepo(s.db)
	}
	return s.event
}

func (s *Store) EventStudent() storage.EventStudentRepoI {
	if s.eventStudent == nil {
		s.eventStudent =NewEventStudentRepo(s.db)
	}
	return s.eventStudent
}

func (s *Store) StudentPayment() storage.StudentPaymentRepoI {
	if s.studentpayment == nil {
		s.studentpayment = NewStudentPaymentRepo(s.db)
	}
	return s.studentpayment
}

