package migrations

import (
	"database/sql"

	"github.com/GuiaBolso/darwin"
)

var items = []darwin.Migration{
	{
		Version:     1,
		Description: "create table",
		Script: `CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`,
	},
}

func Run(db *sql.DB) error {
	driver := darwin.NewGenericDriver(db, darwin.PostgresDialect{})
	migrations := darwin.New(driver, items, nil)
	return migrations.Migrate()
}
