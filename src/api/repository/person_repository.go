package repository

import "teste1/src/api/models"

type PersonRepository interface {
	Save(models.Person) (models.Person, error)
	FindById(uint32) (models.Person, error)
}