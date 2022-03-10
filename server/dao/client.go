/*
Package dao
@Author: MoZhu
@File: client
@Software: GoLand
*/
package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

// Client 使用 gorm 的 Client
type Client interface {
	DB() *gorm.DB
	CtxDB(ctx context.Context) *gorm.DB
}

type ctxTransactionKey struct{}

type MySQLClient struct {
	db *gorm.DB
}

func (c *MySQLClient) DB() *gorm.DB {
	return c.db
}

func (c *MySQLClient) CtxDB(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return c.db.WithContext(ctx)
	}
	iCtxDB := ctx.Value(ctxTransactionKey{})
	ctxDB, ok := iCtxDB.(*gorm.DB)
	if ok {
		return ctxDB
	}
	return c.db.WithContext(ctx)
}

func (c *MySQLClient) CreateObject(module interface{}, value interface{}) error {
	return c.db.Model(module).Create(value).Error
}

func (c *MySQLClient) First(module interface{}) (interface{}, error) {
	var value interface{}
	err := c.db.Model(module).First(value).Error
	return value, err
}

func (c *MySQLClient) FirstOrCreate(module interface{}, where string, args []interface{}, value interface{}) error {
	return c.db.Model(module).Where(where, args...).FirstOrCreate(value).Error
}

func (c *MySQLClient) QueryList(module interface{}, where string, args []interface{}) (int, interface{}, error) {
	// TODO
	return 0, nil, nil
}

func (c *MySQLClient) UpdateObject(module interface{}, value interface{}) error {
	return c.db.Model(module).Updates(value).Error
}

func (c *MySQLClient) UpdateObjectByID(module interface{}, value interface{}, id interface{}) error {
	return c.db.Model(module).Where("id=?", id).Updates(value).Error
}

func (c *MySQLClient) DeleteObject(module interface{}, ids interface{}, field ...string) error {
	if len(field) > 0 {
		return c.db.Delete(module, fmt.Sprintf("%s in (?)", field[0]), ids).Error
	}
	return c.db.Delete(module, "id in (?)", ids).Error
}

func (c *MySQLClient) DeleteObjectByFilter(module interface{}, where string, args []interface{}) error {
	tx := c.db.Begin()
	if tx.Where(where, args...).Delete(module); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return tx.Commit().Error
}
