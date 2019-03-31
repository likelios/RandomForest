package controllers

import (
	query "RandomForest/backend/ch-query"
	"RandomForest/backend/models"
	"database/sql"

	"github.com/astaxie/beego"
)

//Connect - объект подключения к ClickHouse
var Connect *sql.DB

//EventDataController API Для работы с событиями
type EventDataController struct {
	beego.Controller
}

// Get ...
// @Title GetOne
// @Description get Event-Data by id
// @Param	cid		path 	string	true		"The key for staticblock"
// @Param	type		path 	string	true		"Temp Humidity WindSpeed"
// @Success 200
// @Failure 403
// @router /:cid/:type [get]
func (c *EventDataController) Get() {
	customer := c.GetString(":cid")
	typeName := c.GetString(":type")

	var data models.EventDataNice
	data = query.GetEventRows(Connect, typeName, customer)
	c.Data["json"] = data
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}

// Get ...
// @Title GetOne
// @Description get Event-Data by id
// @Param	cid		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403
// @router /grid/:cid [get]
func (c *EventDataController) GetGrid() {
	customer := c.GetString(":cid")

	var data []models.EventNiceGrid
	data = query.GetEventTable(Connect, customer)
	c.Data["json"] = data
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}
