package configs

import (
	"log"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func InitDatabase() {
	err_load_env := godotenv.Load("local.env")

	if err_load_env != nil {
		log.Fatalf("Some error occured. Err: %s", err_load_env)
	}

	DATABASE_TYPE := os.Getenv("DATABASE_TYPE")
	USER_DB := os.Getenv("USER_DB")
	PASS_WORD_DB := os.Getenv("PASS_WORD_DB")
	SCHEMA_DB := os.Getenv("SCHEMA_DB")

	orm.RegisterDriver(DATABASE_TYPE, orm.DRMySQL)
	err_register_orm := orm.RegisterDataBase("default", DATABASE_TYPE, USER_DB+":"+PASS_WORD_DB+"@/"+SCHEMA_DB+"?charset=utf8")

	if err_register_orm != nil {
		glog.Fatal("Failed to register dtatbase %v", err_register_orm)
	}

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err_register_orm = orm.RunSyncdb(name, force, verbose)
	if err_register_orm != nil {
		glog.Fatal("Failed to run sync %v", err_register_orm)
	}
}
