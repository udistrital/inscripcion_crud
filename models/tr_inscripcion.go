package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrInscripcionReintegro struct {
	InscripcionEstudiante *Inscripcion
	ReintegroEstudiante   *Reintegro
}


// AddTrInscripcionReintegro Transacción para registrar toda la información de un reintegro
func AddTrInscripcionReintegro(m *TrInscripcionReintegro) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	
	if idInscripcion, errTr := o.Insert(m.InscripcionEstudiante); errTr == nil {
		fmt.Println("Inscripcion registrada, Id: ", idInscripcion)
		m.ReintegroEstudiante.InscripcionId.Id = int(idInscripcion)
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
		return
	}

	if idReintegro, errTr := o.Insert(m.ReintegroEstudiante); errTr == nil {
		fmt.Println("Reintegro registrado, Id: ", idReintegro)
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
		return
	}

	_ = o.Commit()

	return
}
