package controllers

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	//"teste1/src/api/auth"
	"teste1/src/api/database"
	"teste1/src/api/models"
	"teste1/src/api/repository"
	"teste1/src/api/repository/crud"
	"teste1/src/api/responses"

	"github.com/gorilla/mux"
)

/*
	Func *READ*: Obter dados de uma pessoa 
	TODO: Verificar se quem faz o request é admin
	--->> Extrair o token que é recebido no request e verificar se pode realizar a operação
*/
func GetPerson(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}


	/*token := auth.DecodeToken(w,r)
	fmt.Println("ID DO TOKEN:", token.Claims.(*models.Claim).ID)
	fmt.Println("ID DO REQUEST:", uid)
	if token != nil {
		if token.Claims.(*models.Claim).ID != uint32(uid){
			responses.ERROR(w,http.StatusUnauthorized,errors.New(http.StatusText(http.StatusUnauthorized)))
			return
		}
	}*/

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPersonCRUD(db)

	func(personRepository repository.PersonRepository) {
		person, err := personRepository.FindById(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, person)
	}(repo)

}
/*
	Func *CREATE*: Criar utilizador
	TODO: Verificar se quem faz o request é admin
	--->> Extrair o token que é recebido no request e verificar se pode realizar a operação
*/
func CreatePerson(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	person := models.Person{}
	err = json.Unmarshal(body,&person)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	/*
	token := auth.DecodeToken(w,r)
	fmt.Println("ID DO TOKEN:", token.Claims.(*models.Claim).ID)
	fmt.Println("ID DO REQUEST:", person.ID)
	if token != nil {
		if token.Claims.(*models.Claim).ID == person.ID{
			responses.ERROR(w,http.StatusUnauthorized,errors.New(http.StatusText(http.StatusBadRequest)))
			return
		}
		
	}*/

	//fmt.Println("USER: ", uid)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPersonCRUD(db)

	func(personRepository repository.PersonRepository) {
		person, err := personRepository.Save(person)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, person.ID))
		responses.JSON(w, http.StatusCreated, person)
	}(repo)

}
