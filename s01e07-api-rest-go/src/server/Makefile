include Makefile.inc

db:
	sqlite3 topics.db "CREATE TABLE topic(id INTEGER PRIMARY KEY, text TEXT, like INTEGER);"

deps:
	go get -v github.com/julienschmidt/httprouter
	go get -v github.com/mattn/go-sqlite3

fmt:
	for d in $(DIRS); do \
		go fmt $$d; \
	done
	go fmt main.go

build:
	go build -o rest-api main.go

tests:
	rm -f tests.db
	sqlite3 tests.db "CREATE TABLE topic(id INTEGER PRIMARY KEY, text TEXT, like INTEGER);"
	go test interfaces
