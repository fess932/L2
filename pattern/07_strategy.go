package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Выбор алгоритмов обработки в зависимости от входящих данных ( на лету)
Часто применяется для выбора авторизации пользователя

Например вход по логину и паролю, вход через ВК, Гугл и тд

минусы
 Усложняет программу за счёт дополнительных сущностей.
 Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

*/

func strategyUsage() {
	s := Service{}

	// login via username and password
	s.SetStrategy(NewNamePasswordStrategy("login", "password"))
	s.Login()

	// login via vk
	s.SetStrategy(NewVKStrategy("vk token"))
	s.Login()

	// login via google
	s.SetStrategy(NewGoogleStrategy("google token"))
	s.Login()
}

type LoginStrategy interface {
	Login() bool
}

type Service struct {
	LoginStrategy
}

func (s *Service) SetStrategy(strategy LoginStrategy) {
	s.LoginStrategy = strategy
}

/////////////////////////////////

type NamePasswordStrategy struct {
	name, password string
}

func (nps NamePasswordStrategy) Login() bool {
	fmt.Println("login via username and password", nps.name, nps.password)
	return true
}
func NewNamePasswordStrategy(name, password string) *NamePasswordStrategy {
	return &NamePasswordStrategy{name, password}
}

///////////////////////////////////

type GoogleStrategy struct {
	token string
}

func (nps GoogleStrategy) Login() bool {
	fmt.Println("login via google", nps.token)

	return true
}
func NewGoogleStrategy(token string) *GoogleStrategy {
	return &GoogleStrategy{token}
}

///////////////////////////////////

type VKStrategy struct {
	token string
}

func (nps VKStrategy) Login() bool {
	fmt.Println("login via vk", nps.token)

	return true
}
func NewVKStrategy(token string) *VKStrategy {
	return &VKStrategy{token}
}
