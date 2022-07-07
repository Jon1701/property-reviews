/* Create application user */
DO
$do$
BEGIN
	IF EXISTS (SELECT FROM pg_user WHERE usename='appuser') THEN
		RAISE NOTICE 'Skipping application user creation due to existing user';
	ELSE
		CREATE USER appuser WITH PASSWORD 'appuser';
	END IF;
END
$do$;