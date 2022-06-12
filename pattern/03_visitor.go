package pattern

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Посетитель позволяет применять одну и ту же операцию к различным экземпляром типов без их изменения.

Вместо добавления метода во множество типов
лучше использовать паттерн посетитель.
Создать новый тип с нужным методом для каждого типа
или функцию с обработкой каждого нужного типа
*/

func VisitorUsage() {
	cat := Cat{"Василий"}
	dog := Dog{"Бобик"}
	lion := Lion{Surename: "Левит"}

	log.Println(jsonVisitor(cat))
	log.Println(jsonVisitor(dog))
	log.Println(jsonVisitor(lion))
}

func jsonVisitor(e interface{}) ([]byte, error) {
	switch v := e.(type) {
	case Cat:
		return json.Marshal(map[string]string{"name": v.Name})
	case Dog:
		return json.Marshal(map[string]string{"name": v.Nickname})
	case Lion:
		return json.Marshal(map[string]string{"name": v.Surename})
	default:
		return nil, fmt.Errorf("unknown type")
	}
}

type Cat struct {
	Name string
}

type Dog struct {
	Nickname string
}

type Lion struct {
	Surename string
}
