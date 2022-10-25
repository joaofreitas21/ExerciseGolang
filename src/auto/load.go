package auto

import (
	"log"
	"teste1/src/api/database"
	"teste1/src/api/models"
	"teste1/src/api/utils/console"
)

func Load() {
	db, err := database.Connect()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Person{}).Error
	if err != nil{
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.Person{}).Error
	if err != nil{
		log.Fatal(err)
	}

	for _, person := range persons {
		err = db.Debug().Model(&models.Person{}).Create(&person).Error
		if err != nil{
			log.Fatal(err)
		}
		console.Pretty(person)
	}
}