BEGIN;

	DROP EXTENSION IF EXISTS "pgcrypto";
	DROP EXTENSION IF EXISTS "citext";
	DROP EXTENSION IF EXISTS "uuid-ossp";

	DROP FUNCTION IF EXISTS trigger_updated_at_timestamp() 

COMMIT;
