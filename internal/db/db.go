package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Connection struct {
	*gorm.DB
}

func NewDBConn() (*Connection, error) {
	//Todo: Remove hardcoded values
	dbConString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		"root",
		"password",
		"localhost",
		"3306",
		"crud_api_db",
	)

	db, err := gorm.Open(mysql.Open(dbConString), &gorm.Config{
		//this will accept the singular table names in DB
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &Connection{
		DB: db,
	}, nil

}
