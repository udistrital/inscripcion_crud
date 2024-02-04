package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type DeleteConstraintUqAspirante_20240124_165351 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &DeleteConstraintUqAspirante_20240124_165351{}
	m.Created = "20240124_165351"

	migration.Register("DeleteConstraintUqAspirante_20240124_165351", m)
}

// Run the migrations
func (m *DeleteConstraintUqAspirante_20240124_165351) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE inscripcion.inscripcion DROP CONSTRAINT uq_aspirante")
}

// Reverse the migrations
func (m *DeleteConstraintUqAspirante_20240124_165351) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE inscripcion.inscripcion ADD CONSTRAINT uq_aspirante UNIQUE (persona_id, programa_academico_id, periodo_id);")	
}
