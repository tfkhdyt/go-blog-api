all: start

build:
	go build -o ./bin/blog-api ./cmd/blog-api

install:
	go install ./cmd/blog-api

start: db-start build
	./bin/blog-api $(ARGS)

dev: db-start
	go run ./cmd/blog-api $(ARGS)

# dependency:
# - watchexec = https://github.com/watchexec/watchexec
dev-watch: db-start
	watchexec -c -r -e go -- go run ./cmd/blog-api $(ARGS)

sqlc-generate:
	sqlc generate

test:
	go test -v ./...

test-cover:
	go test -v -cover ./...

test-cover-watch:
	watchexec -c -r -e go -- go test -v -cover ./...

test-cover-html:
	go test -coverprofile cover.out ./... && \
		go tool cover -html=cover.out

generate:
	go generate ./...

clean:
	rm -f ./bin/* ./cover.out

db-start:
	systemctl is-active postgresql || systemctl start postgresql

db-stop:
	systemctl is-active postgresql && systemctl stop postgresql

db-status:
	systemctl status postgresql

.PHONY: build install start dev dev-watch sqlc-generate test test-cover test-cover-watch test-cover-html generate clean db-start db-stop db-status
