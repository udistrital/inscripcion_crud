package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsTipoDocumentoPrograma_20210114_164522 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsTipoDocumentoPrograma_20210114_164522{}
	m.Created = "20210114_164522"

	migration.Register("InsertsTipoDocumentoPrograma_20210114_164522", m)
}

// Run the migrations
func (m *InsertsTipoDocumentoPrograma_20210114_164522) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210114_164522_inserts_tipo_documento_programa.up.sql")
	if err != nil {
		//handle error
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
func (m *InsertsTipoDocumentoPrograma_20210114_164522) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210114_164522_inserts_tipo_documento_programa.down.sql")
	if err != nil {
		//handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
