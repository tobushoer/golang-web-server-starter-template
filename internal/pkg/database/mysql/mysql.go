package mysql

import (
	"fmt"
	"log"
	"template/internal/pkg/config"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// New create singleton mysql connection
func New(cfg *config.Config) *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	mysqlCfg := cfg.MySQL
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True",
		mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Address, mysqlCfg.Port, mysqlCfg.DBName)

	tryTimes := 3
	for i := 0; i < tryTimes; i++ {
		db, err = gorm.Open("mysql", args)
		if err == nil {
			break
		}
		log.Printf("Connect to %s failed: %s\n", args, err.Error())
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Panicf("Connect to %s failed: %s\n", args, err.Error())
	}

	go func() {
		for {
			if err := db.DB().Ping(); err == nil {
				return
			}
			log.Print(err.Error())
			time.Sleep(3 * time.Second)
		}
	}()

	return db
}

// DB Get db instance
func DB() *gorm.DB {
	if db == nil {
		log.Panic("Please new db first")
	}
	return db
}
