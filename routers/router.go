// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/inscripcion_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/inscripcion_pregrado",
			beego.NSInclude(
				&controllers.InscripcionPregradoController{},
			),
		),

		beego.NSNamespace("/inscripcion_posgrado",
			beego.NSInclude(
				&controllers.InscripcionPosgradoController{},
			),
		),

		beego.NSNamespace("/reintegro",
			beego.NSInclude(
				&controllers.ReintegroController{},
			),
		),

		beego.NSNamespace("/tipo_inscripcion",
			beego.NSInclude(
				&controllers.TipoInscripcionController{},
			),
		),

		beego.NSNamespace("/inscripcion",
			beego.NSInclude(
				&controllers.InscripcionController{},
			),
		),

		beego.NSNamespace("/transferencia",
			beego.NSInclude(
				&controllers.TransferenciaController{},
			),
		),

		beego.NSNamespace("/tipo_icfes",
			beego.NSInclude(
				&controllers.TipoIcfesController{},
			),
		),

		beego.NSNamespace("/estado_inscripcion",
			beego.NSInclude(
				&controllers.EstadoInscripcionController{},
			),
		),

		beego.NSNamespace("/propuesta",
			beego.NSInclude(
				&controllers.PropuestaController{},
			),
		),

		beego.NSNamespace("/tipo_proyecto",
			beego.NSInclude(
				&controllers.TipoProyectoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
