package database

import (
	"errors"
	// "os"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

type Db interface {
	Debug() *gorm.DB
	Exec(sql string, values ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Group(query string) *gorm.DB
	Having(query interface{}, values ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	Limit(limit interface{}) *gorm.DB
	Offset(offset interface{}) *gorm.DB
	Order(value interface{}, reorder ...bool) *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Table(name string) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Update(attrs ...interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}

func Connect() (Db, error) {
	dbDriver := "sql"

	switch dbDriver {
	case "sql":
		return connectMySQL()
	}

	return nil, errors.New("database driver error")
}
