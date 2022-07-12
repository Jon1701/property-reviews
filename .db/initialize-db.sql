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
								CHECK (email_address = lower(email_address))
								UNIQUE,

	password			TEXT
								NOT NULL
);
