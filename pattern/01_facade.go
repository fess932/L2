package pattern

import "net/http"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

Скрытие сложности, унификация доступа к сложной системе.
Антипаттерн - становление god object
*/

// Facade - предоставляет простой интерфейс для сложной системы.
// Пример на вводе строки в поисковую строку браузера.
func FacadeSearch() {
	// обработка ввода текста
	searchString := typing()

	// создание http запроса
	req := createHttpRequest(searchString)

	// отправка запроса на сервер
	resp := sendHttpRequest(req)

	// обработка ответа от сервера
	view := processHttpResponse(resp)

	// вывод ответа на экран
	view.Show()
}

func mouseClick(target string) {

}

func typing() string {
	return ""
}

func createHttpRequest(query string) *http.Request {
	return nil
}

func sendHttpRequest(request *http.Request) *http.Response {
	return nil
}

func processHttpResponse(response *http.Response) View {
	return View{}
}

type View struct {
}

func (v View) Show() {

}
