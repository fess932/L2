package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

/*
=== Утилита sort ===
Отсортировать строки (man sort)
Основное
Поддержать ключи
-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
Дополнительное
Поддержать ключи
-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f, err := os.Open("./develop/dev03/task.go")
	if err != nil {
		log.Fatal(err)
	}

	slice := make([]string, 0, 50)

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		slice = append(slice, buf.Text())
	}

	sort.Strings(slice)

	for _, v := range slice {
		os.Stdout.WriteString(v + "\n")
	}
}
