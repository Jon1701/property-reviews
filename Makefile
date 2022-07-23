SERVER_PORT=9000
POSTGRES_DOCKER_CONTAINER_NAME=postgres
POSTGRES_ADMIN_USERNAME=postgres
POSTGRES_ADMIN_PASSWORD=postgres
POSTGRES_DB=property-reviews
POSTGRES_ADMIN_CONNSTRING=postgresql://${POSTGRES_ADMIN_USERNAME}:${POSTGRES_ADMIN_PASSWORD}@localhost/${POSTGRES_DB}

POSTGRES_APP_CONNSTRING=postgresql://${POSTGRES_ADMIN_USERNAME}:${POSTGRES_ADMIN_PASSWORD}@localhost/${POSTGRES_DB}

JWT_SIGNING_KEY=031a9c015be2438bbf9ffbb1e2911038

# Starts services and initializes the database.
start: stop-services start-services wait-5s initialize-db populate-db

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

# Stops all Docker Compose services and deletes all volumes.
destroy:
	@echo "Stopping services and deleting volumes..."
	@docker-compose down --volumes
	@echo "Done stopping services and deleting volumes"

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
	SERVER_PORT=${SERVER_PORT} \
	POSTGRES_APP_CONNSTRING=${POSTGRES_APP_CONNSTRING} \
	JWT_SIGNING_KEY=${JWT_SIGNING_KEY} \
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

# Populates the database with sample data.
populate-db:
	@echo "Loading sample data..."
	@docker container cp ./.db/load-sample-data.sql ${POSTGRES_DOCKER_CONTAINER_NAME}:/tmp
	@docker container exec \
		-t ${POSTGRES_DOCKER_CONTAINER_NAME} \
			psql ${POSTGRES_ADMIN_CONNSTRING} \
				-f /tmp/load-sample-data.sql
	@echo "Done loading sample data"