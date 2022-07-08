SERVER_PORT=9000
POSTGRES_DOCKER_CONTAINER_NAME=postgres
POSTGRES_ADMIN_USERNAME=postgres
POSTGRES_ADMIN_PASSWORD=postgres
POSTGRES_DB=property-reviews
POSTGRES_ADMIN_CONNSTRING=postgresql://${POSTGRES_ADMIN_USERNAME}:${POSTGRES_ADMIN_PASSWORD}@localhost/${POSTGRES_DB}

# Starts services and initializes the database.
start: stop-services start-services wait-5s initialize-db

# Starts Docker Compose services.
start-services:
	@echo "Starting services..."
	POSTGRES_DOCKER_CONTAINER_NAME=${POSTGRES_DOCKER_CONTAINER_NAME} \
	POSTGRES_ADMIN_USERNAME=${POSTGRES_ADMIN_USERNAME} \
	POSTGRES_ADMIN_PASSWORD=${POSTGRES_ADMIN_PASSWORD} \
	POSTGRES_DB=${POSTGRES_DB} \
		docker-compose up -d
	@echo "Done starting services"

# Stops Docker Compose services.
stop-services:
	@echo "Stopping services..."
	@docker-compose down
	@echo "Done stopping services"

# Opens a shell to the Postgres Docker container.
db-shell:
	@echo "Opening a shell to \`${POSTGRES_DOCKER_CONTAINER_NAME}\` Docker container..."
	@docker exec -it ${POSTGRES_DOCKER_CONTAINER_NAME} /bin/bash
	@echo "Done opening shell"

# Executes psql in container.
psql:
	@echo "Executing psql..."
	@docker exec -it ${POSTGRES_DOCKER_CONTAINER_NAME} \
		bash -c "psql ${POSTGRES_ADMIN_CONNSTRING}"
	@echo "Done executing psql"

# Executes main.go.
run:
	@echo "Running main.go"
	@SERVER_PORT=${SERVER_PORT} \
		go run main.go
	@echo "Done running main.go"

# Wait 5 seconds.
wait-5s:
	@echo "Waiting for 5 seconds..."
	@sleep 5s
	@echo "Done waiting for 5 seconds"

# Run the database initialization script.
initialize-db:
	@echo "Initializing database..."
	@docker container cp ./.db/initialize-db.sql ${POSTGRES_DOCKER_CONTAINER_NAME}:/tmp
	@docker container exec \
		-t ${POSTGRES_DOCKER_CONTAINER_NAME} \
			psql ${POSTGRES_ADMIN_CONNSTRING} \
				-f /tmp/initialize-db.sql
	@echo "Done initialization database"