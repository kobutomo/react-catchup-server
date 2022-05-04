package main

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	models "github.com/kobutomo/react-catchup-server/src/infrastructure/dbmodels"
)

func main() {
	db, err := sql.Open("mysql", "root:root@(localhost:3306)/react_catchup?parseTime=true")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	ctx := context.Background()
	dm := models.User{
		Username: "tomoro",
		Email:    "n.tomoro13@gmail.com",
		Password: "Password",
	}
	dm.Insert(ctx, db, boil.Infer())
	m, err := models.Users().One(ctx, db)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	println(m.Username)
}
