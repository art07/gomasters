.SILENT:

# Run app
run:
	go run main.go

# Lint check
lint:
	golangci-lint run

# Mockgen
gen:
	mockgen -source=handler/handler.go -destination=mocks/handler.go -package=mocks
