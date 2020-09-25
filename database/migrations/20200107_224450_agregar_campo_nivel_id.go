package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoNivelId_20200107_224450 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoNivelId_20200107_224450{}
	m.Created = "20200107_224450"

	migration.Register("AgregarCampoNivelId_20200107_224450", m)
}

// Run the migrations
func (m *AgregarCampoNivelId_20200107_224450) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE inscripcion.tipo_inscripcion ADD COLUMN nivel_id integer ;")
}

// Reverse the migrations
func (m *AgregarCampoNivelId_20200107_224450) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE inscripcion.tipo_inscripcion DROP COLUMN  nivel_id ;")
}
