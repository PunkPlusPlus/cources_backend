package storage

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	DB *gorm.DB
}

var instance *Storage
var once sync.Once

func GetStorage() *Storage {
	//once.Do(func() {
	//	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//	db, err := sql.Open("postgres", psqlconn)
	//	if err != nil {
	//		panic(err)
	//	}
	//	//db.SetConnMaxLifetime(time.Minute * 3)
	//	db.SetMaxOpenConns(10)
	//	db.SetMaxIdleConns(10)
	//	err = db.Ping()
	//	if err != nil {
	//		panic(err)
	//	}
	//	instance = &Storage{DB: db}
	//})
	once.Do(func() {
		// https://github.com/go-gorm/postgres
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  "host=localhost user=admin password=admin dbname=warden port=5432 sslmode=disable",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = &Storage{DB: db}
	})
	return instance
}
