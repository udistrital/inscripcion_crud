package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Inscripcion struct {
	Id                  int                `orm:"column(id);pk;auto"`
	PersonaId           int                `orm:"column(persona_id)"`
	ProgramaAcademicoId int                `orm:"column(programa_academico_id);null"`
	ReciboInscripcion   string             `orm:"column(recibo_inscripcion);null"`
	PeriodoId           int                `orm:"column(periodo_id)"`
	EnfasisId           int                `orm:"column(enfasis_id);null"`
	NotaFinal           float64            `orm:"column(nota_final);null"`
	AceptaTerminos      bool               `orm:"column(acepta_terminos)"`
	FechaAceptaTerminos time.Time          `orm:"column(fecha_acepta_terminos);type(date)"`
	Activo              bool               `orm:"column(activo)"`
	FechaCreacion       string             `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion   string             `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	Credencial          int                `orm:"column(credencial);null"`
	Opcion              int                `orm:"column(opcion);null"`
	TipoCupo            int                `orm:"column(tipo_cupo);null"`
	EstadoInscripcionId *EstadoInscripcion `orm:"column(estado_inscripcion_id);rel(fk)"`
	TipoInscripcionId   *TipoInscripcion   `orm:"column(tipo_inscripcion_id);rel(fk)"`
}

func (t *Inscripcion) TableName() string {
	return "inscripcion"
}

func init() {
	orm.RegisterModel(new(Inscripcion))
}

// AddInscripcion insert a new Inscripcion into database and returns
// last inserted Id on success.
func AddInscripcion(m *Inscripcion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInscripcionById retrieves Inscripcion by Id. Returns error if
// Id doesn't exist
func GetInscripcionById(id int) (v *Inscripcion, err error) {
	o := orm.NewOrm()
	v = &Inscripcion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInscripcion retrieves all Inscripcion matches certain condition. Returns empty list if
// no records exist
func GetAllInscripcion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Inscripcion)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.HasSuffix(k, "__in") {
			arr := strings.Split(v, "|")
			qs = qs.Filter(k, arr)
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Inscripcion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateInscripcion updates Inscripcion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInscripcionById(m *Inscripcion) (err error) {
	o := orm.NewOrm()
	v := Inscripcion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInscripcion deletes Inscripcion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInscripcion(id int) (err error) {
	o := orm.NewOrm()
	v := Inscripcion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Inscripcion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
