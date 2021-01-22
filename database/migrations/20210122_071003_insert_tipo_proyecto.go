package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertTipoProyecto_20210122_071003 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertTipoProyecto_20210122_071003{}
	m.Created = "20210122_071003"

	migration.Register("InsertTipoProyecto_20210122_071003", m)
}

// Run the migrations
func (m *InsertTipoProyecto_20210122_071003) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/220210122_071003_insert_tipo_proyecto.up.sql")

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
func (m *InsertTipoProyecto_20210122_071003) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

	file, err := ioutil.ReadFile("../scripts/220210122_071003_insert_tipo_proyecto.down.sql")

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
