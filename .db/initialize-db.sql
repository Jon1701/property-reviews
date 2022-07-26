-- Sets the timestamp.
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

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
								NOT NULL,

	created_at		TIMESTAMP
								NOT NULL
								DEFAULT NOW(),

	updated_at		TIMESTAMP
								NOT NULL
								DEFAULT NOW()							
);

-- Update updated_at when User is updated.
CREATE TRIGGER
	update_user_timestamp_on_update
BEFORE UPDATE ON
	users
FOR EACH ROW
	EXECUTE PROCEDURE trigger_set_timestamp();

-- Create Management Companies table.
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
	
	website							VARCHAR(255),

	created_at					TIMESTAMP
											NOT NULL
											DEFAULT NOW(),

	updated_at					TIMESTAMP
											NOT NULL
											DEFAULT NOW()			
);

-- Create index for the id_hash column in the Management Companies table. 
CREATE INDEX IF NOT EXISTS
	index_management_companies_id_hash
ON
	management_companies(id_hash);

-- Update updated_at when Management Company is updated.
CREATE TRIGGER
	update_management_company_timestamp_on_update
BEFORE UPDATE ON
	management_companies
FOR EACH ROW
	EXECUTE PROCEDURE trigger_set_timestamp();

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

	created_at					TIMESTAMP
											NOT NULL
											DEFAULT NOW(),

	updated_at					TIMESTAMP
											NOT NULL
											DEFAULT NOW(),

	CONSTRAINT fk_management_company
		FOREIGN KEY(management_company_id_hash)
			REFERENCES management_companies(id_hash)	
);

-- Update updated_at when Property is updated.
CREATE TRIGGER
	update_property_timestamp_on_update
BEFORE UPDATE ON
	properties
FOR EACH ROW
	EXECUTE PROCEDURE trigger_set_timestamp();

/* Create Reviews table. */
CREATE TABLE IF NOT EXISTS reviews (
	id									SERIAL
											PRIMARY KEY
											NOT NULL,

	id_hash							VARCHAR(50)
											UNIQUE
											NOT NULL,

	user_id_hash				VARCHAR(50),

	property_id_hash		VARCHAR(50),

	description					TEXT,

	overall_rating			NUMERIC(3, 2),

	management_rating		INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	maintenance_rating	INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	cleanliness_rating	INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	heating_rating			INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	storage_rating			INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	safety_rating				INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	cooling_rating			INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	parking_rating			INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	neighborhood_rating	INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	hotwater_rating			INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	amenities_rating		INT
											NOT NULL
											CHECK(management_rating BETWEEN 0 and 5),

	CONSTRAINT fk_user
	FOREIGN KEY(user_id_hash)
		REFERENCES users(id_hash),

	CONSTRAINT fk_property
	FOREIGN KEY(property_id_hash)
		REFERENCES properties(id_hash)
);