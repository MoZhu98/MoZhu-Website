/*
Package dao
@Author: MoZhu
@File: initialize
@Software: GoLand
*/
package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Table interface {
	TableName() string
}

func initDB() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/website?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Init db mysql failed: %v", err)
		return nil
	}
	return db
}

func NewDBClient() *MySQLClient {
	return &MySQLClient{db: initDB()}
}
