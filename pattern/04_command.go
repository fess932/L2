package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Превращение запросов в обьекты
Самый частый пример использования - кнопки в интерфейсе пользователя.
*/

func commandUsage() {
	send := SendButton{}
	save := SaveButton{}

	runCommand(send)
	runCommand(save)
}

type Clicker interface {
	Click()
}

func runCommand(click Clicker) {
	click.Click()
}

// Типа абстрактный класс
type Button struct{}

func (b *Button) Click() {
	println("действие по умолчанию")
}

type SendButton struct {
	Button
}

func (sb SendButton) Click() {
	println("отправка сообщения")
}

type SaveButton struct {
	Button
}

func (sb SaveButton) Click() {
	println("сохранение сообщения")
}
