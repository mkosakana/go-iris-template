package service

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/joho/godotenv"
	"go-iris-sample/_example-mvc-api/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// InitDb gorp初期化処理
func InitDb() *gorp.DbMap {
	// .env ファイルの読み込み
	loadEnv()

	dbUserName := os.Getenv("DB_USERNAME")
	dbPassWord := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDbName := os.Getenv("DB_DATABASE")
	// SQL ドライバとの接続
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", dbUserName, dbPassWord, dbHost, dbPort, dbDbName))
	if err != nil {
		fmt.Printf("error! can't open sql driver %v", err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "utf8mb4"}}

	// users テーブルの作成
	dbmap.AddTableWithName(model.User{}, "users").
		SetKeys(true, "id").
		SetKeys(false, "name").
		SetKeys(false, "age")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		fmt.Printf("error! %v", err)
	}

	//ログの収集
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-iris-sample:", log.Lmicroseconds))

	return dbmap
}

// .envファイル内の設定を取得する
func loadEnv() {
	// .envファイル全体からDB設定を読み込む
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env ファイルの読み込みに失敗しました: %v", err)
	}
}