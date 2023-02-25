include .env

test:
	go clean -testcache
	go test -v ./...

start-webserver:
	python rest-server/main.py
