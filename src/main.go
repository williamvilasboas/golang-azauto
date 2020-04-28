package main

import (
	api "azauto.com.br/api/v2"
	site "azauto.com.br/site/v1"
	"log"
	"net/http"
	"os"
)

func setEnvironments() {
	log.Print("Set environments")
	wd, err := os.Getwd()
	if err != nil {
		log.Print("Fail to Getwd")
	}
	os.Setenv("VIEW_PATH", wd + "/src/views/")
	log.Print("Finish environments")
}

func main() {
	log.Print("Start app")
	setEnvironments()
	api.Run()
	http.HandleFunc("/", site.HelloWord)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
