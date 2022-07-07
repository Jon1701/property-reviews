SERVER_PORT=9000

# Executes main.go.
run:
	@echo "Running main.go"
	@SERVER_PORT=${SERVER_PORT} \
		go run main.go
	@echo "Done running main.go"