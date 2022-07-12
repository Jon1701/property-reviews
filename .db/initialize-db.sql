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

	email_address	VARCHAR(255)
								NOT NULL
								UNIQUE,

	password			TEXT
								NOT NULL
);
