package dbpostgres

import (
	"database/sql"
	"fmt"
	"log"
	config "sensor/pkg/config/sensor"

	//importando de maneira implicita, quem utiliza é o pacote database/sql
	_ "github.com/lib/pq"
)

func getConnectString(config *config.ConfigDatabase) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

}

func PostgresConnection() (*sql.DB, error) {
	config, erro := config.CarregaConfigDB()
	if erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	db, erro := sql.Open("postgres", getConnectString(config))
	if erro != nil {
		log.Fatal(erro)
		db.Close()
		return nil, erro
	}

	return db, nil

}

