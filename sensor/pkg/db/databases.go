package db

import (
	"database/sql"
	config "sensor/pkg/config"
	dbpostgres "sensor/pkg/db/postgres"

	"go.uber.org/zap"
)

type DataSources struct {
	DB *sql.DB
}

func InitializeDataSources(log *zap.Logger, configDatabase *config.ConfigDatabase, configMigrate *config.ConfigMigrate) (*DataSources, error) {
	/*
	 * Postegres
	 */
	p := dbpostgres.NewPostgres(configDatabase, log)
	db, err := p.PostgresConnection()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	m := dbpostgres.NewMigratePostgres(db, configMigrate, log)
	m.MigratePostgres()

	return &DataSources{DB: db}, nil
}
