package connections

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

// PostgresConn is
func PostgresConn(host, user, password, dbName, appMode string, port, maxOpenConn, maxIdleConn int) (dbConn *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
	dbConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if !strings.Contains(appMode, "release") {
		dbConn = dbConn.Debug()
	}
	dbConfig, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	err = dbConfig.Ping()
	if err != nil {
		return nil, err
	}
	dbConfig.SetConnMaxLifetime(time.Minute * 3)
	dbConfig.SetMaxOpenConns(maxOpenConn)
	dbConfig.SetMaxIdleConns(maxIdleConn)

	log.Println("postgres database connected")
	return dbConn, nil
}
