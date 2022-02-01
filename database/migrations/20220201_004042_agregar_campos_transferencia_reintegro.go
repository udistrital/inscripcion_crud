package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposTransferenciaReintegro_20220201_004042 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposTransferenciaReintegro_20220201_004042{}
	m.Created = "20220201_004042"

	migration.Register("AgregarCamposTransferenciaReintegro_20220201_004042", m)
}

// Run the migrations
func (m *AgregarCamposTransferenciaReintegro_20220201_004042) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20220201_004042_agregar_campos_transferencia_reintegro.up.sql")
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
func (m *AgregarCamposTransferenciaReintegro_20220201_004042) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20220201_004042_agregar_campos_transferencia_reintegro.down.sql")

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
