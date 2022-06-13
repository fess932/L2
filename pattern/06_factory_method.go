package pattern

import (
	"fmt"
	"log"
	"time"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Можно применять чтобы уменьшить повторение кода инициализации
Так же можно менять инициализацию каждого отдельного объекта
Например можно добавить в каждый объект какие-то дополнительные поля не затрагивая
основной интерфейс.

Фабрика которая использует метод обьектов для создания обьектов в едином интерфейсе

Минусы - приходится использовать везде метод New()(god object)

*/

func factoryMethodUsage() {
	entites := []Entity{
		New(WolfType),
		New(WolfType),
		New(GoblinType),
		New(SamuraiType),
		New(GoblinType),
	}

	// некий game loop
	for {
		for _, e := range entites {
			e.Action()
		}

		time.Sleep(time.Second)
	}
}

// New Фабрика
func New(etype EntityType) Entity {
	switch etype {
	case WolfType:
		return NewWolf() // Фабричный метод
	case GoblinType:
		return NewGoblin()
	case SamuraiType:
		return NewSamurai()
	default:
		log.Fatalln("unknown entity type")
		return nil
	}

	return nil
}

type Entity interface {
	Action()
}

type EntityType int

const (
	WolfType EntityType = iota + 1
	GoblinType
	SamuraiType
)

type Wolf struct {
	name string
}

func (w Wolf) Action() {
	fmt.Printf("wolf %s: attack you \n", w.name)
}

func NewWolf() Wolf {
	return Wolf{name: "BadWolf"}
}

type Goblin struct {
	name string
}

func (g Goblin) Action() {
	fmt.Printf("goblin %s: attack you \n", g.name)
}

func NewGoblin() Goblin {
	return Goblin{name: "Gaar"}
}

type Samurai struct {
	name string
}

func (g Samurai) Action() {
	fmt.Printf("samurai %s says: we have city to burn\n", g.name)
}

func NewSamurai() Samurai {
	return Samurai{name: "Johnny"}
}
