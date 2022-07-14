/* Create Users table */
CREATE TABLE IF NOT EXISTS users (
	id						SERIAL
								PRIMARY KEY
								NOT NULL,

	id_hash				VARCHAR(50)
								UNIQUE
								NOT NULL,

	email_address	VARCHAR(255)
								NOT NULL
								CHECK (email_address = lower(email_address))
								UNIQUE,

	password			TEXT
								NOT NULL
);

/* Create Management Companies table */
CREATE TABLE IF NOT EXISTS management_companies (
	id									SERIAL
											PRIMARY KEY
											NOT NULL,

	id_hash							VARCHAR(50)
											UNIQUE
											NOT NULL,

	name								VARCHAR(1000)
											NOT NULL,
	
	address_line1				VARCHAR(1000)
											NOT NULL,

	address_line2				VARCHAR(1000),

	address_city				VARCHAR(1000)
											NOT NULL,

	address_state				VARCHAR(1000)
											NOT NULL,

	address_postal_code	VARCHAR(20)
											NOT NULL,

	address_country			VARCHAR(100)
											NOT NULL,
	
	website							VARCHAR(255)
);

/* Create Properties table */
CREATE TABLE IF NOT EXISTS properties (
	id									SERIAL
											PRIMARY KEY
											NOT NULL,

	id_hash							VARCHAR(50)
											UNIQUE
											NOT NULL,

	management_company_id_hash	VARCHAR(50),

	address_line1				VARCHAR(1000)
											NOT NULL,

	address_line2				VARCHAR(1000),

	address_city				VARCHAR(1000)
											NOT NULL,

	address_state				VARCHAR(1000)
											NOT NULL,

	address_postal_code	VARCHAR(20)
											NOT NULL,

	address_country			VARCHAR(100)
											NOT NULL,
	
	property_type				VARCHAR(50)
											NOT NULL,

	building_type				VARCHAR(50)
											NOT NULL,

	neighborhood				VARCHAR(255),

	CONSTRAINT fk_management_company
		FOREIGN KEY(management_company_id_hash)
			REFERENCES management_companies(id_hash)	
);