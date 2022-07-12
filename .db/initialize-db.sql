-- /* Create application user */
-- DO
-- $do$
-- BEGIN
-- 	IF EXISTS (SELECT FROM pg_user WHERE usename='appuser') THEN
-- 		RAISE NOTICE 'Skipping application user creation due to existing user';
-- 	ELSE
-- 		CREATE USER appuser WITH PASSWORD 'appuser';
-- 	END IF;
-- END
-- $do$;

-- /* Enable module to generate UUIDs */
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/* Create Users table */
CREATE TABLE IF NOT EXISTS users (
	id						SERIAL
								PRIMARY KEY
								NOT NULL,

	id_hash				CHAR(50)
								UNIQUE
								NOT NULL,

	username			VARCHAR(50)
								NOT NULL,

	email_address	VARCHAR(255)
								NOT NULL
								UNIQUE,

	password			TEXT
								NOT NULL
);
