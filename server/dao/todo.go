/*
Package dao
@Author: MoZhu
@File: todo
@Software: GoLand
*/
package dao

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `gorm:"content"`
	Done    bool   `gorm:"done"`
}

func (todo Todo) TableName() string {
	return "todo"
}

func NewTodoDAO(client *MySQLClient) TodoDAOIF {
	return &todoDAO{
		client: client,
	}
}

type TodoDAOIF interface {
}

type todoDAO struct {
	client Client
}
