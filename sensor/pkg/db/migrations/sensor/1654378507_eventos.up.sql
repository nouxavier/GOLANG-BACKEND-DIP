BEGIN;

CREATE TABLE eventos (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    id_sensor uuid  NOT NULL,
    valor VARCHAR  NOT NULL,
    created_at 		timestamptz NOT NULL DEFAULT NOW(),
	updated_at 		timestamptz,
    PRIMARY KEY (id),
    CONSTRAINT fk_sensor FOREIGN KEY (id_sensor) REFERENCES sensores (id)

);

CREATE TRIGGER "tr_eventos_updated_at" 
   BEFORE UPDATE ON "eventos"
   FOR EACH ROW 
   EXECUTE FUNCTION trigger_updated_at_timestamp();


COMMIT;