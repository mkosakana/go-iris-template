package service

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	"go-iris-sample/_example-mvc-api/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func initDb() *gorp.DbMap {
	dbUserName := "iris"
	dbPassWord := "root"
	dbHost := "localhost"
	dbPort := "3306"
	dbDbName := "go_iris_database"
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbUserName, dbPassWord, dbHost, dbPort, dbDbName))
	if err != nil {
		fmt.Printf("error! can't open sql driver %v", err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8-mb4"}}

	dbMap.AddTableWithName(model.User{}, "user").SetKeys(true, "Id")
	err = dbMap.CreateTablesIfNotExists()
	if err != nil {
		fmt.Printf("error! %v", err)
	}
	dbMap.TraceOn("[gorp]", log.New(os.Stdout, "go_iris_sample:", log.Lmicroseconds))

	return dbMap
}