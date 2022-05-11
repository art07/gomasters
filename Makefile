.SILENT:

# Run app
run:
	go run main.go

# Lint check
lint:
	golangci-lint run

# Mockgen
gen:
	mockgen -source=usecase/user/usecase.go -destination=mock/user_repo.go -package=mock
