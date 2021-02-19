package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/dimall-id/lumos/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB () error {
	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=UTC",
		config.GetString("db.host"),
		config.GetString("db.username"),
		config.GetString("db.password"),
		config.GetString("db.database"),
		config.GetString("db.port"))

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return err
	}

	Db = db
	var sqlDB *sql.DB
	sqlDB, err = Db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	//err = Db.AutoMigrate(&model.Product{})
	//if err != nil {
	//	return err
	//}

	return nil
}