package database

import (
    "database/sql"
    "log"
    
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    connStr := "user=postgres password=password123 dbname=taskflow sslmode=disable"
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error conectando a la base de datos:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("No se puede conectar a la base de datos:", err)
    }

    log.Println("✅ Conectado a PostgreSQL")
}

func Close() {
    DB.Close()
}