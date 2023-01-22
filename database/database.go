package database

import (
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Database struct {
	Session *xorm.Engine
}

func Open(dtbs *xorm.Engine) *Database {

	d := new(Database)
	d.Session = dtbs
	d.Session.SetMapper(names.GonicMapper{})
	d.Session.DB().Begin()
	d.Session.Sync(new(GuildUser))

	return d
}
