BEGIN;
-- O pgcrypt o módulo fornece funções criptográficas para PostgreSQL
-- https://x-team.com/blog/storing-secure-passwords-with-postgresql/
	CREATE EXTENSION IF NOT EXISTS "pgcrypto" WITH SCHEMA public;
-- O citextmódulo fornece um tipo de cadeia de caracteres que não 
-- diferencia maiúsculas de minúsculas, citext. 
	CREATE EXTENSION IF NOT EXISTS "citext" WITH SCHEMA public;
--fornece funções para gerar identificadores universalmente exclusivos
-- (UUIDs) usando um dos vários algoritmos padrão.
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

-- Automatically updating a timestamp column in PostgreSQL
	CREATE OR REPLACE FUNCTION trigger_updated_at_timestamp()
    RETURNS TRIGGER AS $$
	BEGIN
    	NEW.updated_at = NOW();
		RETURN NEW;
	END;
	$$ LANGUAGE PLPGSQL;

COMMIT;