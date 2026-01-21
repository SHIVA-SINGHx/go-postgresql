package main

import (
	"fmt"
	"go-postgrsql/router"
	"log"
	"net/http"
)

func main(){
	r:= router.Router()
	fmt.Println("Server startin on port 8083..")

	log.Fatal(http.ListenAndServe(":8083", r))
}