package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTipoInscripcionFaltantes_20200113_150705 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoInscripcionFaltantes_20200113_150705{}
	m.Created = "20200113_150705"

	migration.Register("InsertTipoInscripcionFaltantes_20200113_150705", m)
}

// Run the migrations
func (m *InsertTipoInscripcionFaltantes_20200113_150705) Up() {
	m.SQL("INSERT INTO inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Nueva', 'Inscripción de un aspirante nuevo nivel pregrado', 'NUEPRE', true, 1, 14, now(), now()); ")
	m.SQL("INSERT INTO  inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Generación E', 'Inscripción de un aspirante que pertence al programa generación E', 'GENE', true, 8, 14, now(), now()); ")
	m.SQL("INSERT INTO  inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Nueva', 'Inscripción de un aspirante nuevo nivel posgrado', 'NUEPOS', true, 2, 15, now(), now()); ")

}

// Reverse the migrations
func (m *InsertTipoInscripcionFaltantes_20200113_150705) Down() {
	m.SQL("delete from inscripcion.tipo_inscripcion;")

}
