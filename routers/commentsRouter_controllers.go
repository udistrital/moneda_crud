package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TipoMonedaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"] = append(beego.GlobalControllerRouter["github.com/udistrital/moneda_crud/controllers:TrmController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
