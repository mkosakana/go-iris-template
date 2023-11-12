package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mkosakana/go-iris-template/_example-mvc-api/configs/db/schemes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const GormDnsNameFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

// InitDb DB初期化
func InitDb() (*gorm.DB, error) {
	db, err := getConfiguredGorm()
	if err != nil {
		return nil, err
	}

	err = migrateDb(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// gorm設定
func getConfiguredGorm() (*gorm.DB, error) {
	err := godotenv.Load("./environment/.env")
	if err != nil {
		return nil, err
	}

	// https://github.com/go-gorm/mysql
	dnsName := fmt.Sprintf(GormDnsNameFormat,
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_CONNECTION"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dnsName,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

var booksDummyRecords = []schemes.Book{
	{Title: "コンビニ人間", Author: "村田 沙耶香"},
	{Title: "老人と海", Author: "アーネスト・ヘミングウェイ"},
	{Title: "歯車", Author: "芥川 龍之介"},
}

// migration
func migrateDb(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if !db.Migrator().HasTable(&schemes.Book{}) {
			if err := db.AutoMigrate(&schemes.Book{}); err != nil {
				return err
			}
			if err := db.Create(&booksDummyRecords).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
