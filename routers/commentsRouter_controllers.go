package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:OrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:RelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoOrganizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/organizacion_crud/controllers:TipoRelacionOrganizacionesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
