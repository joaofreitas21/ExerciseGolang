package middlewares

import (
	//"context"
	//"fmt"
	"log"
	"net/http"
	"teste1/src/api/auth"
	//"teste1/src/api/responses"
	//"teste1/src/api/models"
	//"teste1/src/api/utils/types"
)

// SetMiddlewareLogger displays a info message of the API
func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

// SetMiddlewareJSON set the application Content-Type
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//func to authorize an access
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		//token := auth.DecodeToken(r)
		//console.Pretty(token)
		//next(w, r)

		token := auth.DecodeToken(w, r)
		if token == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, r)
		

	}
}


