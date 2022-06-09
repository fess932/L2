package main

import (
	"l2/develop/dev11/errors"
	"log"
	"net/http"
)

/*
=== HTTP server ===
Реализовать HTTP сервер для работы с календарем.
В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API,
       используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов

Методы API:
POST /create_event
POST /update_event
POST/delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ
содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
       В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
       В случае остальных ошибок сервер должен возвращать HTTP 500.
       Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	r := NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.ListenAndServe(":8080", r)
}

type GRouter struct {
	routes map[string]route
}

type route struct {
	get  func(http.ResponseWriter, *http.Request)
	post func(http.ResponseWriter, *http.Request)
}

func (gr *GRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	log.Println(r.Method)
	log.Println(r.RequestURI)

	if _, ok := gr.routes[r.URL.Path]; !ok {
		errors.JSONError(w, http.StatusNotFound, errors.ErrNotFound)
	}
}

func NewRouter() *GRouter {
	return &GRouter{
		routes: make(map[string]route),
	}
}

func (gr *GRouter) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	var rt = gr.routes[pattern]
	rt.get = handler
	gr.routes[pattern] = rt
}

func (gr *GRouter) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	var rt = gr.routes[pattern]
	rt.post = handler
	gr.routes[pattern] = rt
}

//
//func newRouter() {
//
//}
//
//func router() {
//
//}
//
//func router() {
//
//	mux := http.NewServeMux()
//	mux.Handle("/", http.FileServer(http.Dir("static")))
//
//	c := &CalendarServce{}
//	mux.Handle("/create_event", c)
//
//	http.HandleFunc("/", mux.ServeHTTP)
//
//	//POST /create_event
//	//POST /update_event
//	//POST/delete_event
//	//GET /events_for_day
//	//GET /events_for_week
//	//GET /events_for_month
//}
//
//type CalendarServce struct {
//}
//
//func (c *CalendarServce) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//}
