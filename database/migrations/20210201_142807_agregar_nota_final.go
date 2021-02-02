package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarNotaFinal_20210201_142807 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarNotaFinal_20210201_142807{}
	m.Created = "20210201_142807"

	migration.Register("AgregarNotaFinal_20210201_142807", m)
}

// Run the migrations
func (m *AgregarNotaFinal_20210201_142807) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210201_142807_agregar_nota_final.up.sql")

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
func (m *AgregarNotaFinal_20210201_142807) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210201_142807_agregar_nota_final.down.sql")

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
