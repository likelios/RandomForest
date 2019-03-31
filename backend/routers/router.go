// @APIVersion 1.0.0
// @Title Случайны лес API
// @Description API подготовленное командой Случайный лес
package routers

import (
	"RandomForest/backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/event",
			beego.NSInclude(
				&controllers.EventDataController{},
			),
		),
		beego.NSNamespace("/client",
			beego.NSInclude(
				&controllers.ClientController{},
			),
		),
		beego.NSNamespace("/company",
			beego.NSInclude(
				&controllers.CompaniController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
