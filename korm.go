package krom

import (
	"database/sql"
	"fmt"
	"korm/dialect"
	"korm/log"
	"korm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
	}
	// Send a ping to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return nil, fmt.Errorf("dialect %s Not Found", driver)
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return e, err
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
