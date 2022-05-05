package mysql

import (
	"database/sql"
	"os"
)

type IMySQL interface {
	GetConn() *sql.DB
}

type MySQL struct {
	conn *sql.DB
}

func NewMySQL() IMySQL {
	c, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/react-cathup?parseTime=true")
	if err != nil {
		os.Exit(1)
	}
	return MySQL{
		conn: c,
	}
}

func (m MySQL) GetConn() *sql.DB {
	return m.conn
}
