package database

import (
	"errors"
	"fmt"
	// "os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func connectMySQL() (*gorm.DB, error) {
	username := "root"
	password := "root"
	host := "mysql"
	port := "3306"
	connection := "sample"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8", username, password, host, port, connection)

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		return nil, errors.New("database connect error")
	}

	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(10)

	return db, nil
}