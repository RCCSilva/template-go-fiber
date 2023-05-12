package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

type Database struct {
	connString string
	DB         *sql.DB
}

func New(connStr string) *Database {
	return &Database{DB: nil, connString: connStr}
}

func (d *Database) Connect() error {
	if d.DB != nil {
		return nil
	}

	db, err := sql.Open("postgres", d.connString)
	if err != nil {
		return err
	}

	d.DB = db

	return nil
}

func (d *Database) Migrate() error {
	err := d.Connect()
	if err != nil {
		return err
	}

	driver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, d.connString)
	if err != nil {
		return err
	}

	if err = m.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}
