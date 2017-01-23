package models

import (
	"fmt"

	"github.com/Ascendings/gowis/modules/settings"

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

	dbSettings = settings.Cfg.Section("database")
)

// set up our DB
func init() {
	// register database drivers
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// register our database connection
	orm.RegisterDataBase("default", dbSettings.Key("driver").String(), createConnectionString())
}

func createConnectionString() string {
	dbDriver := dbSettings.Key("driver").String()
	// check the driver
	if dbDriver == "sqlite3" {
		// we can return the sqlite db path
		return dbSettings.Key("path").String()
	} else if dbDriver == "mysql" || dbDriver == "postgres" {
		// now we need to build a connection string
		return fmt.Sprintf("%s:%s@%s/%s?charset=%s",
			dbSettings.Key("username").String(), dbSettings.Key("password").String(),
			dbSettings.Key("host").String(), dbSettings.Key("db_name").String(),
			dbSettings.Key("charset").String())
	}

	// return nothing here
	return ""
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
