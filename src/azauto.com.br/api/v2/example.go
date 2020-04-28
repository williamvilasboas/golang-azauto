package v2

import (
	"fmt"
	"net/http"
)

func irineu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello word")
}

func Run() {
	http.HandleFunc("/irineu", irineu)
}