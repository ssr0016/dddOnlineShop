package database

import (
	"fmt"
	"onlineShop/internal/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgres(Cfg config.DBConfig) (db *sqlx.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Cfg.Host, Cfg.Port, Cfg.User, Cfg.Password, Cfg.Name)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(Cfg.ConnectionPool.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(Cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(Cfg.ConnectionPool.MaxOpenConnetcion))
	db.SetMaxIdleConns(int(Cfg.ConnectionPool.MaxIdleConnection))

	return
}
