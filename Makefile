.PHONY:build run clean build_and_run migrate

build:
	go build -v ./cmd/api

run:
	./api

clean:
	rm -f ./api

migrate:
	migrate -path migrations -database "postgres://test_user:test_pwd@localhost/test_db?sslmode=disable" up

build_and_run: clean migrate build run

.DEFAULT_GOAL := build_and_run