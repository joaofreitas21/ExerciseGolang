package database

import (
	//"fmt"
	"teste1/src/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB,error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err!= nil{
		//fmt.Println("Erro aqui")
		return nil, err	
	}
	return db, nil
}