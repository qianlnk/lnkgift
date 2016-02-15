/**
 * FileName:		mysql.go
 * Description:		get mysql connect
 * Author:			Qianno.Xie
 * Email:			qianlnk@163.com
**/
package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"haostudent/config"
	"sync"
)

type Handler func(context.Context, *sqlx.DB, interface{}) error

var (
	once sync.Once
	db   *sqlx.DB
)

func Invoke(ctx context.Context, handle Handler, dest interface{}) error {
	return handle(ctx, db, dest)
}

func init() {
	once.Do(initializing)
}

func initializing() {
	conn, err := sqlx.Connect("mysql", config.GetConfig().GetString("database.mysql"))
	if err != nil {
		fmt.Printf("mysql connect err: %+v\n", err)
	}
	db = conn
}

func GetDB() *sqlx.DB {
	return db
}
