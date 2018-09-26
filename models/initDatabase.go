package models

import (
	"../utils"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func RegisterDatabase(driverName string, dataSource string, maxIdle int, maxOpen int) {
	orm.Debug = true
	utils.CheckError(orm.RegisterDriver("postgres", orm.DRPostgres))
	utils.CheckError(orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxOpen))

	registerModels()
	utils.CheckError(orm.RunSyncdb("default", false, true))
}

func registerModels() {

}
