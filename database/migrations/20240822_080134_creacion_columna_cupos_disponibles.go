package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionColumnaCuposDisponibles_20240822_080134 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionColumnaCuposDisponibles_20240822_080134{}
	m.Created = "20240822_080134"

	migration.Register("CreacionColumnaCuposDisponibles_20240822_080134", m)
}

// Run the migrations
func (m *CreacionColumnaCuposDisponibles_20240822_080134) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE IF EXISTS inscripcion.cupo_inscripcion ADD COLUMN cupos_disponibles integer;")
}

// Reverse the migrations
func (m *CreacionColumnaCuposDisponibles_20240822_080134) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE IF EXISTS inscripcion.cupo_inscripcion DROP COLUMN cupos_disponibles;")
}
