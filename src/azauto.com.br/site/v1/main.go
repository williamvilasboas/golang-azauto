package v1

import (
	"log"
	"html/template"
	"net/http"
	"os"
)

type InterfaceSimple struct {
	PageTitle string
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	pathView := os.Getenv("VIEW_PATH")

	t, err := template.ParseFiles(pathView + "homepage.html")

	if err != nil {
		log.Fatal("Template parsing error:", err)
	}

	err = t.Execute(w, InterfaceSimple{
		PageTitle: "Simples teste brother!",
	})
	if err != nil {
		log.Fatal("template executing error: ", err)
	}
}