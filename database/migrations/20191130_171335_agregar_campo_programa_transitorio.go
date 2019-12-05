package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCampoProgramaTransitorio_20191130_171335 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCampoProgramaTransitorio_20191130_171335{}
	m.Created = "20191130_171335"

	migration.Register("AgregarCampoProgramaTransitorio_20191130_171335", m)
}

// Run the migrations
func (m *AgregarCampoProgramaTransitorio_20191130_171335) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191130_171335_agregar_campo_programa_transitorio.up.sql")

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
func (m *AgregarCampoProgramaTransitorio_20191130_171335) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191130_171335_agregar_campo_programa_transitorio.down.sql")

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
