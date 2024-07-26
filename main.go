package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/json", NewCoderjSON)
	http.ListenAndServe(":8080", nil) // Запуск сервера
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет мир!")
}

func NewCoderjSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mymessage := map[string]string{"Name": "Semen"}
	json.NewEncoder(w).Encode(mymessage)
}
