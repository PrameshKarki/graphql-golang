package configs

import (
	"database/sql"
	"log"
	"sync"

	"github.com/PrameshKarki/event-management-golang/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var (
	once sync.Once
	db   *sql.DB
)

// Create a map of table names
var TABLE_NAME = map[string]string{
	"EVENT":       "events",
	"USER":        "users",
	"USER_EVENTS": "user_events",
	"SESSION":     "event_sessions",
	"EXPENSE":     "expenses",
}

func initializeDatabase() {
	cfg := mysql.Config{
		User:                 utils.GoDotEnv("DB_USER"),
		Passwd:               utils.GoDotEnv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               utils.GoDotEnv("DB_NAME"),
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	logrus.Info("Database Connected!")
}

func GetDatabaseConnection() *sql.DB {
	once.Do(initializeDatabase)
	return db
}

func GetDialect() goqu.DialectWrapper {
	return goqu.Dialect("mysql")
}
