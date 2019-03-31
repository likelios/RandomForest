package controllers

import (
	"RandomForest/backend/mysql"

	"github.com/astaxie/beego"
)

//CompaniController API Для работы с компаниями
type CompaniController struct {
	beego.Controller
}

// GetOne ...
// @Title Get One
// @Description get Compani by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Compani
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CompaniController) GetOne() {
	l := mysql.GetCompany(MyConnect, "")
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
// @Description get Compani
// @Success 200 {object} []models.Compani
// @Failure 403
// @router / [get]
func (c *CompaniController) GetAll() {
	l := mysql.GetCompany(MyConnect, "")
	c.Data["json"] = l
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.ServeJSON()
}
