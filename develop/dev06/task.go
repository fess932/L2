package main

import (
	"bufio"
	"flag"
	"fmt"
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
	fields := make([]int, 0, len(raw))

	var (
		err            error
		strFrom, strTo string
		from, to       int
		ok             bool
	)

	for _, field := range raw {
		strFrom, strTo, ok = strings.Cut(field, "-")
		if !ok {
			from, err = strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("invalid field: %w", err)
			}

			fields = append(fields, from)

			continue
		}

		if strFrom != "" {
			from, err = strconv.Atoi(strFrom)
			if err != nil {
				return nil, fmt.Errorf("invalid field: %w", err)
			}
		}

		if from == 0 {
			from = 1
		}

		if strTo == "" {
			to = len(raw)
		} else {
			to, err = strconv.Atoi(strTo)
			if err != nil {
				return nil, fmt.Errorf("invalid field: %w", err)
			}
		}

		for ; from <= to; from++ {
			fields = append(fields, from)
		}
		log.Println("")
	}

	return fields, nil
}

func cut(r io.Reader, w io.Writer, fields []int, delimiter string, separated bool) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		columns := strings.Split(s.Text(), delimiter)
		if separated && len(columns) < 2 {
			continue
		}

		for i, field := range fields {
			if field > len(columns) {
				continue
			}

			fmt.Fprint(w, columns[field-1])
			if i < len(fields)-1 {
				fmt.Fprint(w, delimiter)
			}
		}
		fmt.Fprintln(w)
	}
}
