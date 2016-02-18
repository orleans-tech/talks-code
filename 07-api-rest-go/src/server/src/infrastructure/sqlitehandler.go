package infrastructure

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqliteHandler struct {
	Path    string
	Connect *sql.DB
}

func (handler *SqliteHandler) Open() error {
	db, err := sql.Open("sqlite3", handler.Path)
	if err != nil {
		log.Panic(err)
	}
	handler.Connect = db
	return err
}

func (handler *SqliteHandler) Close() {
	if handler.Connect != nil {
		handler.Connect.Close()
		handler.Connect = nil
	}
}

func (handler SqliteHandler) Db() *sql.DB {
	return handler.Connect
}

func NewSqliteHandler(path string) *SqliteHandler {
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Path = path
	return sqliteHandler
}
