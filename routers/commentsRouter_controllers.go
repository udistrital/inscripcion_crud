package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:EstadoInscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPosgradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:InscripcionPregradoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:PropuestaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:ReintegroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoIcfesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoInscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TipoProyectoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TrInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TrInscripcionController"],
        beego.ControllerComments{
            Method: "PostReintegro",
            Router: `/reintegro`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TrInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TrInscripcionController"],
        beego.ControllerComments{
            Method: "PostTransferencia",
            Router: `/transferencia`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/inscripcion_crud/controllers:TransferenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
