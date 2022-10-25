package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"teste1/src/api/auth"
	"teste1/src/api/models"
	"teste1/src/api/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	token, err := auth.SignIn(person.Username, person.Password)
	if err != nil {
		responses.ERROR(w, http.StatusForbidden, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}
