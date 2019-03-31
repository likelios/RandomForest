package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"],
        beego.ControllerComments{
            Method: "GetByCompany",
            Router: `/getbycompany/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:ClientController"],
        beego.ControllerComments{
            Method: "GetByNotCompany",
            Router: `/getbynotcompany/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:CompaniController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:CompaniController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:CompaniController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:CompaniController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:EventDataController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:EventDataController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:cid/:type`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:EventDataController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:EventDataController"],
        beego.ControllerComments{
            Method: "GetGrid",
            Router: `/grid/:cid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["RandomForest/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["RandomForest/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
