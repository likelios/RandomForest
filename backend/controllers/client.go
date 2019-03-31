package controllers

import (
	"RandomForest/backend/mysql"

	"github.com/astaxie/beego"
)

//  API для работы с клиентом
type ClientController struct {
	beego.Controller
}

// GetOne ...
// @Title Get One
// @Description get Client by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Client
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ClientController) GetOne() {
	customer := c.GetString(":id")
	l := mysql.GetClient(MyConnect, customer)
	if len(l) > 0 {
		c.Data["json"] = l[0]
	}
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Client
// @Success 200 {object} []models.Client
// @Failure 403
// @router / [get]
func (c *ClientController) GetAll() {
	l := mysql.GetClient(MyConnect, "")
	c.Data["json"] = l
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Client by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} []models.Client
// @Failure 403 :id is empty
// @router /getbycompany/:id [get]
func (c *ClientController) GetByCompany() {
	customer := c.GetString(":id")

	c.Data["json"] = mysql.GetClientByCompany(MyConnect, customer, " = ")

	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Client by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} []models.NiceClient
// @Failure 403 :id is empty
// @router /getbynotcompany/:id [get]
func (c *ClientController) GetByNotCompany() {
	customer := c.GetString(":id")

	c.Data["json"] = mysql.GetClientByCompany(MyConnect, customer, " != ")

	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}
