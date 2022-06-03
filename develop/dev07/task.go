package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===
Реализовать функцию, которая будет объединять один или более done каналов
в single канал если один из его составляющих каналов закроется.

Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая,
приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}
start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)
fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})

		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or2(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(2*time.Second),
		sig(15*time.Millisecond),
		sig(15*time.Millisecond),
		sig(15*time.Millisecond),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v\n", time.Since(start))
}

// 100% CPU usage
func or(channels ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer close(ch)

		for {
			var active bool

			for _, c := range channels {
				select {
				case <-c:
					active = true
				default:
				}
			}

			if active {
				return
			}
		}
	}()

	return ch
}

// or2 запускать на каждый канал горутину и как только
// одна из горутин закроется, закрывать все горутины и делать возврат
func or2(channels ...<-chan interface{}) <-chan interface{} {
	orChan := make(chan interface{})

	for _, v := range channels {
		go func(c <-chan interface{}) {
			select {
			case <-orChan: // срабатывает на закрытие канала
				return

			case <-c:
				close(orChan)

				return
			}
		}(v)
	}

	return orChan
}
