package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoTipoCupo_20240801_230708 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoTipoCupo_20240801_230708{}
	m.Created = "20240801_230708"

	migration.Register("AgregarCampoTipoCupo_20240801_230708", m)
}

// Run the migrations
func (m *AgregarCampoTipoCupo_20240801_230708) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE IF EXISTS inscripcion.inscripcion ADD COLUMN tipo_cupo integer;")
	m.SQL("ALTER TABLE IF EXISTS inscripcion.documento_programa ADD COLUMN tipo_cupo integer;")

}

// Reverse the migrations
func (m *AgregarCampoTipoCupo_20240801_230708) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE IF EXISTS inscripcion.inscripcion DROP COLUMN tipo_cupo;")
	m.SQL("ALTER TABLE IF EXISTS inscripcion.documento_programa DROP COLUMN tipo_cupo;")

}
