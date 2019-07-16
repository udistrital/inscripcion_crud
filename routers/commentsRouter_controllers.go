package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
