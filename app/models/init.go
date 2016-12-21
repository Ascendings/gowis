package models

import (
	"fmt"

	"gogs.ballantine.tech/gballan1/gowis/modules/settings"

	"github.com/astaxie/beego/orm"
	// loading SQLite dialect for beego ORM
	_ "github.com/mattn/go-sqlite3"
	// loading MySQL dialect for beego ORM
	_ "github.com/go-sql-driver/mysql"
	// loading Postgres dialect for beego ORM
	_ "github.com/lib/pq"
)

var (
	// DB - beego database engine object
	DB orm.Ormer
)

// set up our DB
func init() {
	// register database drivers
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// register our database connection
	orm.RegisterDataBase("default", settings.Cfg.Section("database").Key("driver").String(), "gowis.db")
}

// InitDB - creates DB connection
func InitDB() {
	// sync our models with our DB
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

	// create new database connection
	DB = orm.NewOrm()
	// set the database connection to use
	DB.Using("default")
}
