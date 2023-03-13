package pgstore

import (
	"context"
	"database/sql"
	"github.com/audetv/fct-parser/repos/question"
	"github.com/google/uuid"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // Postgresql driver
)

var _ question.Store = &Questions{}

type DBPgQuestion struct {
	ID           uuid.UUID `db:"id"`
	DataID       int       `db:"data_id"`
	ParentDataID int       `db:"parent_data_id"`
	Position     int       `db:"position"`
	Username     string    `db:"username"`
	UserRole     string    `db:"user_role"`
	Text         string    `db:"text"`
	Date         time.Time `db:"date"`
}

type Questions struct {
	db *sql.DB
}

func NewQuestions(dsn string) (*Questions, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	qs := &Questions{
		db: db,
	}
	return qs, nil
}

func (us *Questions) Close() {
	us.db.Close()
}

func (us *Questions) Create(ctx context.Context, q question.Question) (*uuid.UUID, error) {
	dbq := &DBPgQuestion{
		ID:           q.ID,
		DataID:       q.DataID,
		ParentDataID: q.ParentDataID,
		Position:     q.Position,
		Username:     q.Username,
		UserRole:     q.UserRole,
		Text:         q.Text,
		Date:         q.Date,
	}

	_, err := us.db.ExecContext(ctx, `INSERT INTO question_comments 
	(id, data_id, question_data_id, position, username, user_role, text, date)
	values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		dbq.ID,
		dbq.DataID,
		dbq.ParentDataID,
		dbq.Position,
		dbq.Username,
		dbq.UserRole,
		dbq.Text,
		dbq.Date,
	)
	if err != nil {
		return nil, err
	}

	return &q.ID, nil
}
