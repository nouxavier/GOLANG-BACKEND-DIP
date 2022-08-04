package repository

const(
	qEvento_Get string = `select id, id_sensor, valor from eventos where id=$1`
	qEvento_GetEventosToIDSensor string = `select id, id_sensor, valor FROM eventos  WHERE id_sensor = $1`
	qEvento_Create string = `insert into eventos (id_sensor, valor) values ($1, $2) returning id`
	
)