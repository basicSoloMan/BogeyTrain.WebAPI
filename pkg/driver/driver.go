package driver

import (
	"database/sql"
	"fmt"

	"BogeyTrain/pkg/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
