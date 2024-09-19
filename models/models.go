package models

import (
	"flash_sale/pkg/setting"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type Model struct {
	ID         uint `gorm:"primary_key" json:"id"`
	CreatedOn  int  `json:"created_on"`
	ModifiedOn int  `json:"modified_on"`
	DeletedOn  int  `json:"deleted_on"`
}

func Setup() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.Callback().Create().Before("gorm:create").Register("update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Before("gorm:delete").Register("delete_callback", deleteCallback)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		now := time.Now().Unix()
		db.Statement.SetColumn("CreatedOn", now)
		db.Statement.SetColumn("ModifiedOn", now)
		db.Statement.SetColumn("DeletedOn", 0)
	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		now := time.Now().Unix()
		db.Statement.SetColumn("ModifiedOn", now)
	}
}

func deleteCallback(db *gorm.DB) {}
