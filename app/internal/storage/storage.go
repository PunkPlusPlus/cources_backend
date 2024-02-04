package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "warden"
)

type Storage struct {
	DB *sql.DB
}

var instance *Storage
var once sync.Once

func GetStorage() *Storage {
	once.Do(func() {
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			panic(err)
		}
		//db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		instance = &Storage{DB: db}
	})
	return instance
}
