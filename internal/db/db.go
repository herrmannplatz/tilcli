package db

import (
	"database/sql"
	"os"
	"path/filepath"
	"time"

	"github.com/herrmannplatz/tilcli/internal/util"
	_ "github.com/mattn/go-sqlite3"
)

type learning struct {
	ID          uint
	Title       string
	Description string
	CreatedAt   time.Time
}

type LearningDB struct {
	DB      *sql.DB
	DataDir string
}

func (t *LearningDB) CreateTableIfNotExists() error {
	_, err := t.DB.Exec(`
		CREATE TABLE IF NOT EXISTS "learnings" (
			"id" INTEGER,
			"title" TEXT NOT NULL,
			"description" TEXT NOT NULL,
			"createdAt" DATETIME, 
			PRIMARY KEY("id" AUTOINCREMENT)
	)`)
	return err
}

func (l *LearningDB) Insert(title, description string) error {
	_, err := l.DB.Exec(
		"INSERT INTO learnings(title, description, createdAt) VALUES(?, ?, ?)",
		title,
		description,
		time.Now())
	return err
}

func (l *LearningDB) GetAll() ([]learning, error) {
	rows, err := l.DB.Query("SELECT * FROM learnings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	learnings := []learning{}
	for rows.Next() {
		var learning learning
		err := rows.Scan(&learning.ID, &learning.Title, &learning.Description, &learning.CreatedAt)
		if err != nil {
			return nil, err
		}
		learnings = append(learnings, learning)
	}
	return learnings, nil
}

func (l *LearningDB) Random() (learning, error) {
	row := l.DB.QueryRow("SELECT * FROM learnings ORDER BY RANDOM() LIMIT 1")
	var learning learning
	err := row.Scan(&learning.ID, &learning.Title, &learning.Description, &learning.CreatedAt)
	return learning, err
}

func (l *LearningDB) Reset() error {
	if err := os.RemoveAll(l.DataDir); err != nil {
		return err
	}
	return nil
}

func Connect() (*LearningDB, error) {
	path, err := util.GetApplicationDataDir()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", filepath.Join(path, "til.db"))
	if err != nil {
		return nil, err
	}
	l := LearningDB{db, path}
	if err := l.CreateTableIfNotExists(); err != nil {
		return nil, err
	}
	return &l, nil
}
