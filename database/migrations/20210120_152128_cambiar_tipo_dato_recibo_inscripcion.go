package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambiarTipoDatoReciboInscripcion_20210120_152128 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambiarTipoDatoReciboInscripcion_20210120_152128{}
	m.Created = "20210120_152128"

	migration.Register("CambiarTipoDatoReciboInscripcion_20210120_152128", m)
}

// Run the migrations
func (m *CambiarTipoDatoReciboInscripcion_20210120_152128) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210120_152128_cambiar_tipo_dato_recibo_inscripcion.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CambiarTipoDatoReciboInscripcion_20210120_152128) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210120_152128_cambiar_tipo_dato_recibo_inscripcion.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
