package repository

const (
	qSensor_Get    string = `select id, nome from sensores where id=$1`
	qSensor_Create string = `insert into sensores (nome, nome_regiao, nome_pais) values ($1, $2, $3) returning id`
)
