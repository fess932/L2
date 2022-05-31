package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

// Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
// Использовать библиотеку github.com/beevik/ntp.
// Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
//
// Требования:
// Программа должна быть оформлена как go module
// Программа должна корректно обрабатывать ошибки библиотеки:
// выводить их в STDERR и возвращать ненулевой код выхода в OS

func main() {
	t, err := ntp.Time("0.ru.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

		return
	}

	fmt.Println(t)
	fmt.Println(time.Now())
}
