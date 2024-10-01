package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddActivoToSoporteDocumentoPrograma_20241001_161909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddActivoToSoporteDocumentoPrograma_20241001_161909{}
	m.Created = "20241001_161909"

	migration.Register("AddActivoToSoporteDocumentoPrograma_20241001_161909", m)
}

// Run the migrations
func (m *AddActivoToSoporteDocumentoPrograma_20241001_161909) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	// Paso 1: Agregar la columna con un valor predeterminado temporal
	m.SQL("ALTER TABLE inscripcion.soporte_documento_programa ADD COLUMN activo BOOLEAN DEFAULT TRUE;")

	// Paso 2: Actualizar los registros existentes
	m.SQL("UPDATE inscripcion.soporte_documento_programa SET activo = TRUE WHERE activo IS NULL;")

	// Paso 3: Modificar la columna para que sea NOT NULL con el valor predeterminado correcto
	m.SQL("ALTER TABLE inscripcion.soporte_documento_programa ALTER COLUMN activo SET NOT NULL;")
	m.SQL("ALTER TABLE inscripcion.soporte_documento_programa ALTER COLUMN activo SET DEFAULT TRUE;")
}

// Reverse the migrations
func (m *AddActivoToSoporteDocumentoPrograma_20241001_161909) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE inscripcion.soporte_documento_programa DROP COLUMN activo;")
}
