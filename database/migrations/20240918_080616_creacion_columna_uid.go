package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionColumnaUid_20240918_080616 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionColumnaUid_20240918_080616{}
	m.Created = "20240918_080616"

	migration.Register("CreacionColumnaUid_20240918_080616", m)
}

// Run the migrations
func (m *CreacionColumnaUid_20240918_080616) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE IF EXISTS inscripcion.inscripcion_pregrado ADD COLUMN uid integer;")
}

// Reverse the migrations
func (m *CreacionColumnaUid_20240918_080616) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE IF EXISTS inscripcion.inscripcion_pregrado DROP COLUMN uid;")
}
