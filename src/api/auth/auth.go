package auth

import (
	"teste1/src/api/database"
	"teste1/src/api/models"
	"teste1/src/api/security"
	"teste1/src/api/utils/channels"

	"github.com/jinzhu/gorm"
)

func SignIn(username, password string) (string, error) {
	person := models.Person{}
	var err error
	var db *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		//procura o utilizador
		err = db.Debug().Model(models.Person{}).Where("username = ?", username).Take(&person).Error
		if err != nil {
			ch <- false
			return
		}
		//verifica a hash da pw
		err = security.VerifyPassword(person.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		person.Password = ""
		return GenerateToken(person.ID)
	}

	return "", err
}



