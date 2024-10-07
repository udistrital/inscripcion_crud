package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UpdateIdiomaColumnToNullable_20241007_073329 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UpdateIdiomaColumnToNullable_20241007_073329{}
	m.Created = "20241007_073329"

	migration.Register("UpdateIdiomaColumnToNullable_20241007_073329", m)
}

// Run the migrations
func (m *UpdateIdiomaColumnToNullable_20241007_073329) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE inscripcion.inscripcion_posgrado ALTER COLUMN idioma DROP NOT NULL;")
}

// Reverse the migrations
func (m *UpdateIdiomaColumnToNullable_20241007_073329) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE inscripcion.inscripcion_posgrado ALTER COLUMN idioma SET NOT NULL;")
}
