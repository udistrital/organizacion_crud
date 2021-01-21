// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/organizacion_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_organizacion",
			beego.NSInclude(
				&controllers.TipoOrganizacionController{},
			),
		),

		beego.NSNamespace("/tipo_relacion_organizaciones",
			beego.NSInclude(
				&controllers.TipoRelacionOrganizacionesController{},
			),
		),

		beego.NSNamespace("/relacion_organizaciones",
			beego.NSInclude(
				&controllers.RelacionOrganizacionesController{},
			),
		),

		beego.NSNamespace("/organizacion",
			beego.NSInclude(
				&controllers.OrganizacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
