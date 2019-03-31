package main

import (
	"RandomForest/backend/controllers"
	_ "RandomForest/backend/routers"
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/kshvakov/clickhouse"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	controllers.Connect = newCHConnect()
	controllers.MyConnect = newMySQL()

	beego.Run()
}

func newCHConnect() *sql.DB {
	db, err := sql.Open("clickhouse", "tcp://116.203.112.27:9000?debug=false")
	if err != nil {
		log.Println(err)
	}
	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
	}

	return db
}

func newMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:11223344@tcp(116.203.112.27:3306)/rfteamdb?parseTime=true")
	if err != nil {
		log.Println(err)
	}

	return db
}
