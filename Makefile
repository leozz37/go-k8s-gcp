NAME=medium

## build: Compile the packages
build:
	@go build -o $(NAME)

## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development

## deps: Download modules
deps:
	@go mod download

##prepare-env: Creates .env file with environmnets variable
prepare-env:
	@echo 'DATABASE_HOST=""\n\
	DATABASE_USER="root"\n\
	DATABASE_PASS="toor"'> .env

## test: Run tests with verbose mode
test: prepare-env
	@go test -v -covermode=count -coverprofile=coverage.out ./...
