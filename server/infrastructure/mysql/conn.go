package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type IMySQL interface {
	GetConn() *sql.DB
}

type MySQL struct {
	conn *sql.DB
}

func NewMySQL() IMySQL {
	c, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/react_catchup?parseTime=true")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	if err := c.Ping(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return MySQL{
		conn: c,
	}
}

func (m MySQL) GetConn() *sql.DB {
	return m.conn
}
