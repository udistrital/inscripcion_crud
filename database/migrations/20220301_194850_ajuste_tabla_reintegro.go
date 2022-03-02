package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteTablaReintegro_20220301_194850 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteTablaReintegro_20220301_194850{}
	m.Created = "20220301_194850"

	migration.Register("AjusteTablaReintegro_20220301_194850", m)
}

// Run the migrations
func (m *AjusteTablaReintegro_20220301_194850) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220301_194850_ajuste_tabla_reintegro.up.sql")
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
func (m *AjusteTablaReintegro_20220301_194850) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220301_194850_ajuste_tabla_reintegro.down.sql")

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
