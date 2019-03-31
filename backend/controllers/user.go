package controllers

import (
	"database/sql"
	"encoding/json"
	"RandomForest/backend/models"
	"RandomForest/backend/mysql"

	"github.com/astaxie/beego"
)

//Connect - объект подключения к MySQL
var MyConnect *sql.DB

//  UserController operations for User
type UserController struct {
	beego.Controller
}

// Login ...
// @Title Post
// @Description create User
// @Param	body		body 	models.Login	true		"body for User content"
// @Success 201 {object} models.LoginGoodResponse
// @Failure 403 body is empty
// @router /login [post]
func (c *UserController) Login() {
	var ob models.Login
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

	good, err := mysql.GetLogin(MyConnect, ob.Login, ob.PasswordHash)

	if good.ID == -1 {
		c.Data["json"] = err
	} else {
		c.Data["json"] = good
	}
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	c.ServeJSON()
}
