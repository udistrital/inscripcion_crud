package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertEstadosInscripcionAdmision_20200210_135821 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertEstadosInscripcionAdmision_20200210_135821{}
	m.Created = "20200210_135821"

	migration.Register("InsertEstadosInscripcionAdmision_20200210_135821", m)
}

// Run the migrations
func (m *InsertEstadosInscripcionAdmision_20200210_135821) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20200210_135821_insert_estados_inscripcion_admision.up.sql")

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
func (m *InsertEstadosInscripcionAdmision_20200210_135821) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20200210_135821_insert_estados_inscripcion_admision.down.sql")

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
