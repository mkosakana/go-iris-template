package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const TestGormDnsNameFormat = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

// GetTestGormDb テスト用DB取得
func GetTestGormDb(pathToEnvFile string) (*gorm.DB, error) {
	return testInitDb(pathToEnvFile)
}

// テスト用DB初期化
func testInitDb(pathToEnvFile string) (*gorm.DB, error) {
	db, err := testGetConfiguredGorm(pathToEnvFile)
	if err != nil {
		return nil, err
	}

	err = migrateDb(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// テスト用gorm設定
func testGetConfiguredGorm(pathToEnvFile string) (*gorm.DB, error) {
	err := godotenv.Load(pathToEnvFile)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-gorm/mysql
	dnsName := fmt.Sprintf(TestGormDnsNameFormat,
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("TEST_MYSQL_HOST"),
		os.Getenv("TEST_MYSQL_PORT"),
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
