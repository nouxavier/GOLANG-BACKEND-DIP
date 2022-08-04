#!/bin/bash
set -e

if [ -n "$POSTGRES_DBS" ]; then
	for db in $(echo $POSTGRES_DBS | tr ',' ' '); do
		psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<- EOSQL
			CREATE DATABASE $db ENCODING 'UTF8';
		EOSQL
	done
fi