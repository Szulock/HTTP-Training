package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/json", NewCoderjSON)
	http.HandleFunc("/headerget", GetHeaders)
	http.HandleFunc("/getallinfo", getinfo)
	http.HandleFunc("/hello", name)
	http.HandleFunc("/image", imageHandler)

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

func getinfo(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	url := r.URL.String()

	headers := r.Header

	// Формируем ответ в текстовом формате
	response := fmt.Sprintf("Method: %s\nURL: %s\nHeaders:\n", method, url)
	for key, values := range headers {
		for _, value := range values {
			response += fmt.Sprintf("%s: %s\n", key, value)
		}
	}

	// Устанавливаем заголовок Content-Type как text/plain
	w.Header().Set("Content-Type", "text/plain")

	// Отправляем ответ
	w.Write([]byte(response))
}

// Создайте HTTP-сервер, который отвечает "Hello, [name]!" на запросы к /hello?name=[name].
func name(w http.ResponseWriter, r *http.Request) {

	// Получаем значение параметра "name" из URL запроса
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World" // Значение по умолчанию, если параметр "name" не указан
	}
	// Формируем ответ
	response := fmt.Sprintf("Hello, %s!", name)
	// Пишем ответ в ResponseWriter
	w.Write([]byte(response))
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	// Определяем путь к изображению
	imagePath := filepath.Join("png-transparent-heart-pump-to-download-blood-red-donation-red-cross-life-donor-compress-thumbnail.png")

	// Открываем файл изображения
	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Image not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "image/png")

	// Копируем содержимое файла в ResponseWriter
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Failed to serve image.", http.StatusInternalServerError)
	}
}
