package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// GetDB return an instance of postgres gorm db
func GetDB(host string, port string, user string, password string, dbName string, sslMode string) (*gorm.DB, error) {
	return gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbName, password, sslMode))
}
