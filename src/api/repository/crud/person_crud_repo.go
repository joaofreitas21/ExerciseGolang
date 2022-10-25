package crud

import (
	"errors"
	"teste1/src/api/models"
	"teste1/src/api/utils/channels"

	"github.com/jinzhu/gorm"
)

type personCRUD struct {
	db *gorm.DB
}

func NewRepositoryPersonCRUD(db *gorm.DB) *personCRUD{
	return &personCRUD{db}
}

func (r *personCRUD) Save(person models.Person) (models.Person,error){
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Person{}).Create(&person).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return person, nil
	}
	return models.Person{}, err

}

func (r *personCRUD) FindById(uid uint32) (models.Person,error){
	var err error
	person := models.Person{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Person{}).Where("id = ?", uid).Take(&person).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return person, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Person{}, errors.New("User Not Found")
	}
	return models.Person{}, err

}

