package db

import (
	"douyin_backend_microService/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init :init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(Follow{})
	if err != nil {
		panic(err)
	}

	//if err = DB.Use(gormopentracing.New()); err != nil {
	//	panic(err)
	//}
}
