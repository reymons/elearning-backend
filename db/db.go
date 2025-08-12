package db

import (
    "errors"
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Initialize(url string) error {
    if DB != nil {
        return errors.New("initialize db: an attempt to re-initialize database")
    }
    db, err := sql.Open("pgx", url) 
    DB = db
    return err
}

func Close() {
    if DB != nil {
        DB.Close()
    }
}

