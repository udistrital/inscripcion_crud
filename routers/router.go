// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/admisiones_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/admision",
			beego.NSInclude(
				&controllers.AdmisionController{},
			),
		),

		beego.NSNamespace("/enfasis",
			beego.NSInclude(
				&controllers.EnfasisController{},
			),
		),

		beego.NSNamespace("/estado_admision",
			beego.NSInclude(
				&controllers.EstadoAdmisionController{},
			),
		),

		beego.NSNamespace("/grupo_investigacion",
			beego.NSInclude(
				&controllers.GrupoInvestigacionController{},
			),
		),

		beego.NSNamespace("/linea_investigacion",
			beego.NSInclude(
				&controllers.LineaInvestigacionController{},
			),
		),

		beego.NSNamespace("/periodo_academico",
			beego.NSInclude(
				&controllers.PeriodoAcademicoController{},
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

		beego.NSNamespace("/entrevista",
			beego.NSInclude(
				&controllers.EntrevistaController{},
			),
		),

		beego.NSNamespace("/entrevistador",
			beego.NSInclude(
				&controllers.EntrevistadorController{},
			),
		),

		beego.NSNamespace("/estado_entrevista",
			beego.NSInclude(
				&controllers.EstadoEntrevistaController{},
			),

		),
		beego.NSNamespace("/criterio_admision",
			beego.NSInclude(
				&controllers.CriterioAdmisionController{},
			),
		),
		beego.NSNamespace("/tipo_criterio",
			beego.NSInclude(
				&controllers.TipoCriterioController{},
			),
		),
		
	)
	beego.AddNamespace(ns)
}
