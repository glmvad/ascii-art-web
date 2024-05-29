package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Обработчик POST-запроса, который принимает текст, изменяет его и возвращает обратно
func submitHandler(w http.ResponseWriter, r *http.Request) {

	// Получаем текст из запроса
	text := r.FormValue("text")
	format := r.FormValue("format")


	// Читаем стандартный файл
	doubleCh, err := ReadFile(format)
	if err != nil {
		http.Error(w, fmt.Sprintf("ошибка чтения стандартного файла: %v", err), http.StatusInternalServerError)
		return
	}

	// Создаем буфер для записи результата
	// var result strings.Builder

	arg := strings.Replace(text, "\\n", "\n", -1)
	slicearg := strings.Split(arg, "\n")
	for _, val := range arg {
		if val != 10 && val < 32 || val > 126 {
			fmt.Fprint(w, "wrong symbols in argument")
			return
		}
	}
	for i, text := range slicearg {
		if text == "" {
			if !CheckisEmpty(slicearg) {
				fmt.Println()
			} else if i != 0 && CheckisEmpty(slicearg) {
				fmt.Println()
			}
		} else {
			for k := 1; k < 9; k++ {
				for _, val := range text {
					fmt.Fprint(w, doubleCh[(val-32)*9+rune(k)])

				}
				fmt.Fprint(w, "\n")

			}
		}
	}
	// Отправляем результат обратно клиенту

}

func main() {
	// Регистрируем обработчики запросов
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/ascii-art", submitHandler)
	http.ListenAndServe(":2030", nil)
}
func ReadFile(ds string) ([]string, error) {
	standard, err := os.ReadFile(ds)
	if err != nil {
		return nil, err
	}
	chars := strings.Split(strings.ReplaceAll(string(standard), "\r", ""), "\n")
	return chars, err
}
func CheckisEmpty(arStr []string) bool {
	for i := range arStr {
		if arStr[i] != "" {
			return false
		}
	}
	return true
}
