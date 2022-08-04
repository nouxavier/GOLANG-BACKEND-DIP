// package dbpostgres https://pkg.go.dev/github.com/golang-migrate/migrate
package dbpostgres

import (
	"database/sql"
	config "sensor/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

type MigratePostgres struct {
	db     *sql.DB
	config *config.ConfigMigrate
	log    *zap.Logger
}

func NewMigratePostgres(db *sql.DB, config *config.ConfigMigrate, log *zap.Logger) *MigratePostgres {
	return &MigratePostgres{db, config, log}
}

func (mp *MigratePostgres) MigratePostgres() {
	if !mp.config.MigrateStart {
		return
	}
	driver, err := postgres.WithInstance(mp.db, &postgres.Config{})
	/*
	 NewWithDatabaseInstance returns a new Migrate instance from a source URL
	 and an existing database instance. The source URL scheme is defined by each driver.
	 Use any string that can serve as an identifier during logging as databaseName.
	 You are responsible for closing the underlying database client if necessary.
	*/
	if err != nil {
		mp.log.Sugar().DPanic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		mp.config.Path,
		"postgres",
		driver,
	)

	m.Up()
}
