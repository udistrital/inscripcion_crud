package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertCampoNivelId_20200107_225459 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertCampoNivelId_20200107_225459{}
	m.Created = "20200107_225459"

	migration.Register("InsertCampoNivelId_20200107_225459", m)
}

// Run the migrations
func (m *InsertCampoNivelId_20200107_225459) Up() {
	m.SQL("DELETE FROM inscripcion.tipo_inscripcion WHERE codigo_abreviacion = 'PREG';")
	m.SQL("DELETE FROM inscripcion.tipo_inscripcion WHERE codigo_abreviacion = 'POSG';")
	m.SQL("UPDATE inscripcion.tipo_inscripcion SET nivel_id= 14 WHERE codigo_abreviacion = 'PROPED';")
	m.SQL("UPDATE inscripcion.tipo_inscripcion SET nivel_id= 14 WHERE codigo_abreviacion = 'TRANSINT';")
	m.SQL("UPDATE inscripcion.tipo_inscripcion SET nivel_id= 14 WHERE codigo_abreviacion = 'TRANSEXT';")
	m.SQL("UPDATE inscripcion.tipo_inscripcion SET nivel_id= 14 WHERE codigo_abreviacion = 'REING';")
	m.SQL("UPDATE inscripcion.tipo_inscripcion SET nivel_id= 14 WHERE codigo_abreviacion = 'MOVACA';")
	m.SQL("UPDATE inscripcion.tipo_inscripcion SET nivel_id= 14 WHERE codigo_abreviacion = 'PROFETEC';")
	m.SQL("INSERT INTO inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Transferencia interna', 'Inscripción de una transferencia interna', 'TRANSINT', true, 4, 15, now(), now()); ")
	m.SQL("INSERT INTO  inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Transferencia externa', 'Inscripción de una transferencia externa', 'TRANSEXT', true, 5, 15, now(), now()); ")
	m.SQL("INSERT INTO  inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Reingreso', 'Inscripción de un reingreso', 'REING', true, 6, 15, now(), now()); ")
	m.SQL("INSERT INTO  inscripcion.tipo_inscripcion (nombre, descripcion, codigo_abreviacion, activo, numero_orden, nivel_id, fecha_creacion, fecha_modificacion) VALUES ('Movilidad Académica', 'Inscripción de una movilidad académica', 'MOVACA', true, 7, 15, now(), now());")
}

// Reverse the migrations
func (m *InsertCampoNivelId_20200107_225459) Down() {
	m.SQL("delete from inscripcion.tipo_inscripcion;")

}
