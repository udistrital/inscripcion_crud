package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionColumnaCredencialInscripcion_20240506_113102 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionColumnaCredencialInscripcion_20240506_113102{}
	m.Created = "20240506_113102"

	migration.Register("CreacionColumnaCredencialInscripcion_20240506_113102", m)
}

// Run the migrations
func (m *CreacionColumnaCredencialInscripcion_20240506_113102) Up() {
	m.SQL("ALTER TABLE inscripcion.inscripcion ADD COLUMN credencial INTEGER")
}

// Reverse the migrations
func (m *CreacionColumnaCredencialInscripcion_20240506_113102) Down() {
	m.SQL("ALTER TABLE inscripcion.inscripcion DROP COLUMN credencial")
}
