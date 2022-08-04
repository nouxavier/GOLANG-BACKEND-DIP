BEGIN;

CREATE TABLE sensores (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    nome citext NOT NULL,
    nome_regiao int NOT NULL,
    nome_pais int NOT NULL,
    created_at 		timestamptz NOT NULL DEFAULT NOW(),
	updated_at 		timestamptz,
    PRIMARY KEY (id)
);

   CREATE TRIGGER "tr_sensores_updated_at" 
   BEFORE UPDATE ON "sensores"
   FOR EACH ROW 
   EXECUTE FUNCTION trigger_updated_at_timestamp();


COMMIT;