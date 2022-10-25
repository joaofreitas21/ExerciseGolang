package auto

import "teste1/src/api/models"

//Objetivo: Adicionar um utilizador ao iniciar o programa, role ADMIN
var persons = []models.Person{
	models.Person{Age:20,Username:"admin",Password:"admin",Family:"family",Role:"admin"},
}