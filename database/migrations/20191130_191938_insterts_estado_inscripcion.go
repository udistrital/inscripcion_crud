package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InstertsEstadoInscripcion_20191130_191938 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InstertsEstadoInscripcion_20191130_191938{}
	m.Created = "20191130_191938"

	migration.Register("InstertsEstadoInscripcion_20191130_191938", m)
}

// Run the migrations
func (m *InstertsEstadoInscripcion_20191130_191938) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191130_191938_insterts_estado_inscripcion.up.sql")

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
func (m *InstertsEstadoInscripcion_20191130_191938) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191130_191938_insterts_estado_inscripcion.down.sql")

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
