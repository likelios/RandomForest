package controllers

import (
	"database/sql"

	"github.com/astaxie/beego"
)

//BaseController - Основной контроллер, для наследования
type BaseController struct {
	beego.Controller
}

//MyConnect - объект подключения к MySQL
var MyConnect *sql.DB

//Connect - объект подключения к ClickHouse
var Connect *sql.DB
