package models

import (
	"github.com/astaxie/beego/orm"
)

type TipoEnte struct {
	Id                int     `orm:"column(id);pk;auto"`
	Nombre            string  `orm:"column(nombre)"`
	Descripcion       string  `orm:"column(descripcion);null"`
	CodigoAbreviacion string  `orm:"column(codigo_abreviacion);null"`
	Activo            bool    `orm:"column(activo)"`
	NumeroOrden       float64 `orm:"column(numero_orden);null"`
	FechaModificacion string  `orm:"column(fecha_modificacion);null"`
}

func (t *TipoEnte) TableName() string {
	return "tipo_ente"
}

func init() {
	orm.RegisterModel(new(TipoEnte))
}
