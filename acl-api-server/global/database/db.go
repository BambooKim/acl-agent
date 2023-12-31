package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(datasource *Datasource) {
	dsnFormat := "%s:%s@tcp(%s)/ACL_LIST?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnFormat, datasource.Username, datasource.Password, datasource.Address)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("database connection error")
	}

	DB = db
}

type Datasource struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
