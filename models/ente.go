package models

import (
	"github.com/astaxie/beego/orm"
)

type Ente struct {
	Id       int       `orm:"column(id);pk;auto"`
	TipoEnte *TipoEnte `orm:"column(tipo_ente);rel(fk)"`
}

func (t *Ente) TableName() string {
	return "ente"
}

func init() {
	orm.RegisterModel(new(Ente))
}
