clean: stop
	@docker compose down -v

format:
	@go fmt ./...

init:
	@go get ./...

prep: format tidy vet

start:
	@docker compose up --build

stop:
	@docker compose down

tidy:
	@go mod tidy

vet:
	@go vet ./...
