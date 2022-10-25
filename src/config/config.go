package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT = 0
	SECRETKEY []byte
	DBDRIVER = ""
	DBURL = ""
)

func Load() {
	var err error
	//err = godotenv.Load(".env")
	path_dir := "C:/Users/USER/go/src/teste1/src"
	err = godotenv.Load(filepath.Join(path_dir, ".env"))
	if err != nil {
		//fmt.Println("Passou aqui")
		log.Fatal("Error loading .env file")
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil{
		PORT = 9000
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	//fmt.Println(DBDRIVER)
	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	
	
	//fmt.Println(DBDRIVER)
	//fmt.Println(DBURL)
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}