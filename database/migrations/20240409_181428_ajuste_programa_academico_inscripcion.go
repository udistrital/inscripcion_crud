package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteProgramaAcademicoInscripcion_20240409_181428 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteProgramaAcademicoInscripcion_20240409_181428{}
	m.Created = "20240409_181428"

	migration.Register("AjusteProgramaAcademicoInscripcion_20240409_181428", m)
}

// Run the migrations
func (m *AjusteProgramaAcademicoInscripcion_20240409_181428) Up() {
	m.SQL("ALTER TABLE inscripcion.inscripcion ALTER COLUMN programa_academico_id DROP NOT NULL")

}

// Reverse the migrations
func (m *AjusteProgramaAcademicoInscripcion_20240409_181428) Down() {
	m.SQL("ALTER TABLE inscripcion.inscripcion ALTER COLUMN programa_academico_id SET NOT NULL")
}
