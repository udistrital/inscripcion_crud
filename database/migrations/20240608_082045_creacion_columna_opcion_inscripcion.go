package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionColumnaOpcionInscripcion_20240608_082045 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionColumnaOpcionInscripcion_20240608_082045{}
	m.Created = "20240608_082045"

	migration.Register("CreacionColumnaOpcionInscripcion_20240608_082045", m)
}

// Run the migrations
func (m *CreacionColumnaOpcionInscripcion_20240608_082045) Up() {
	m.SQL("ALTER TABLE inscripcion.inscripcion ADD COLUMN opcion INTEGER")
}

// Reverse the migrations
func (m *CreacionColumnaOpcionInscripcion_20240608_082045) Down() {
	m.SQL("ALTER TABLE inscripcion.inscripcion DROP COLUMN opcion")

}
