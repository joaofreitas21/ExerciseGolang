package api

import (
	"fmt"
	"log"
	"net/http"
	"teste1/src/api/router"
	"teste1/src/auto"
	"teste1/src/config"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening on %d...",config.PORT)
	Listen(config.PORT)
	 
}

func Listen(port int){
	r:= router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port),r ))
}