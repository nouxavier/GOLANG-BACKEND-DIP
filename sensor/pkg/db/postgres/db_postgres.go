package dbpostgres

import (
	"database/sql"
	"fmt"
	config "sensor/pkg/config"

	//importando de maneira implicita, quem utiliza é o pacote database/sql
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Postgres struct {
	log    *zap.Logger
	config *config.ConfigDatabase
}

func NewPostgres(config *config.ConfigDatabase, log *zap.Logger) *Postgres {
	return &Postgres{log, config}
}

func (p *Postgres) PostgresConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", getConnectString(p.config))
	if err != nil {
		p.log.Error(err.Error())
		db.Close()
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil

}

func getConnectString(config *config.ConfigDatabase) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

}
