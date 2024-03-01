package repository

import (
	"os"

	"github.com/go-pg/pg/v10"
	"golang.org/x/sync/errgroup"
)

type Migrator interface {
	Migrate(tables ...Tabler) error
}

type MigratorPostgre struct {
	DB     *pg.DB
	sqlGen SQLgen
}

func NewPostgreDB() *pg.DB {

	return pg.Connect(&pg.Options{
		Addr:     "db:" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
}

func NewMigratorPostgre(db *pg.DB, sqlGen SQLgen) *MigratorPostgre {
	return &MigratorPostgre{
		DB:     db,
		sqlGen: sqlGen,
	}
}

func (m *MigratorPostgre) Migrate(tables ...Tabler) error {
	var errGr errgroup.Group
	for _, table := range tables {
		request := m.sqlGen.CreateTableSQL(table)
		errGr.Go(func() error {
			_, err := m.DB.Exec(request)
			return err
		})

	}

	return errGr.Wait()
}
