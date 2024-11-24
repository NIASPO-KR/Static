package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db      *sql.DB
	Builder squirrel.StatementBuilderType
}

func New(driver, url string) (*Postgres, error) {
	db, err := sql.Open(driver, url)
	if err != nil {
		return nil, fmt.Errorf("sql open %s: %w", driver, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}

	return &Postgres{
		db:      db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (p *Postgres) SqlDB() *sql.DB {
	return p.db
}
