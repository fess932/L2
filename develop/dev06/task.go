package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f := flag.String("f", "", "fields to cut, requered. example: -f 1,2,3")
	d := flag.String("d", "\t", "delimiter, default is TAB")
	s := flag.Bool("s", false, "separated - show only lines with delimiter")

	flag.Parse()

	if *f == "" {
		log.Fatalln("fields is required")
	}

	fields, err := parseFields(*f)
	if err != nil {
		log.Fatalln(err)
	}

	cut(os.Stdin, os.Stdout, fields, *d, *s)
}

// parseField parse field number from string
// -3 > 1,2,3
// 2- > 2,3
// 1-2 > 1,2,3
// 1,2 > 1,2
func parseFields(str string) ([]int, error) {
	raw := strings.Split(str, ",")
	fields := make([]int, len(raw))

	var err error
	for i, field := range raw {
		fields[i], err = strconv.Atoi(field)
		if err != nil {
			return nil, err
		}

		// absolute field number
		if fields[i] < 0 {
			fields[i] = -fields[i]
		}
	}

	return fields, nil
}

func cut(r io.Reader, w io.Writer, fields []int, delimiter string, separated bool) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		columns := strings.Split(s.Text(), delimiter)
		for _, field := range fields {
			if field > len(columns) {
				log.Println("Fields must be less than columns") // todo: remove this line
				continue
			}

		}

	}
}
