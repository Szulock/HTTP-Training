package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/json", NewCoderjSON)
	http.HandleFunc("/headerget", GetHeaders)
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

func GetHeaders(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	headersMap := make(map[string][]string)
	for key, value := range headers {
		headersMap[key] = value
	}
	w.Header().Set("Content-Type", "application/json")

	// Кодируем заголовки в JSON и отправляем в ответ
	json.NewEncoder(w).Encode(headersMap)
}
