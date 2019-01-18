package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:AdmisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:EstadoAdmisionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:GrupoInvestigacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:LineaInvestigacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PeriodoAcademicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/admisiones_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
