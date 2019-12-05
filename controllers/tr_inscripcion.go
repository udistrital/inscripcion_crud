package controllers

import (
	"encoding/json"

	"github.com/udistrital/inscripcion_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// TrInscripcionController operations for TipoTercero
type TrInscripcionController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrInscripcionController) URLMapping() {
	c.Mapping("PostReintegro", c.PostReintegro)
	c.Mapping("PostTransferencia", c.PostTransferencia)
}

// PostReintegro ...
// @Title PostReintegro
// @Description create TrInscripcionReintegro
// @Param	body		body 	models.TrInscripcionReintegro	true		"body for TrInscripcionReintegro content"
// @Success 201 {int} models.TrInscripcionReintegro
// @Failure 400 the request contains incorrect syntax
// @router /reintegro [post]
func (c *TrInscripcionController) PostReintegro() {
	var v models.TrInscripcionReintegro
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddTrInscripcionReintegro(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// PostTransferencia ...
// @Title PostTransferencia
// @Description create TrInscripcionTransferencia
// @Param	body		body 	models.TrInscripcionTransferencia	true		"body for TrInscripcionTransferencia content"
// @Success 201 {int} models.TrInscripcionTransferencia
// @Failure 400 the request contains incorrect syntax
// @router /transferencia [post]
func (c *TrInscripcionController) PostTransferencia() {
	var v models.TrInscripcionTransferencia
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddTrInscripcionTransferencia(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}
