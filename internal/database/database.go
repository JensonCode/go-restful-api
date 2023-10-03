package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB;

func ConnectPlanetScale() (*sql.DB) {

    db, err := sql.Open("mysql", os.Getenv("LOCAL_DEV_DSN"))
    
    if err != nil {
        log.Fatalf("failed to load %v", err)
        return nil
    }
    
    err = db.Ping();
    if err != nil{
        log.Fatalf("failed to ping: %v", err)
        return nil
    }
    
    DB = db

    log.Println("Connected to PlanetScale!")

    return db
}
