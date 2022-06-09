package main

import (
	"fmt"
	"l2/develop/dev11/configs"
	"l2/develop/dev11/errors"
	"l2/develop/dev11/pkg"
	"log"
	"net/http"
	"time"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := configs.NewConfig()

	cal := pkg.NewCalendar()
	cd := pkg.NewCalendarHTTPDelivery(cal)
	r := NewRouter()

	logMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("[%s][%s] %q %v\n", r.Method, r.RemoteAddr, r.URL.String(), time.Since(t))
		})
	}

	r.Use(logMiddleware)

	r.Get("/events_for_day", cd.GetEventForDay)
	r.Get("/events_for_week", nil)
	r.Get("/events_for_month", nil)

	r.Post("/create_event", cd.CreateEvent)
	r.Post("/update_event", nil)
	r.Post("/delete_event", nil)

	log.Println("server listening at", config.Addr)
	log.Println(http.ListenAndServe(config.Addr, r))
}

func NewRouter() *GRouter {
	return &GRouter{
		routes: make(map[string]route),
	}
}

// GRouter is a simple router implementation.
type GRouter struct {
	handler     http.Handler
	middlewares []func(http.Handler) http.Handler
	routes      map[string]route
}
type route struct {
	get  func(http.ResponseWriter, *http.Request)
	post func(http.ResponseWriter, *http.Request)
}

// Use add middleware to router
func (gr *GRouter) Use(middlewares ...func(http.Handler) http.Handler) {
	if gr.middlewares == nil {
		gr.middlewares = middlewares

		return
	}

	gr.middlewares = append(gr.middlewares, middlewares...)
}
func (gr *GRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// init handler at first request
	if gr.handler == nil {
		gr.handler = http.HandlerFunc(gr.routeHTTP)
		for _, v := range gr.middlewares {
			gr.handler = v(gr.handler)
		}
	}

	gr.handler.ServeHTTP(w, r)
}
func (gr *GRouter) routeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, ok := gr.routes[r.URL.Path]; !ok {
		errors.JSONError(w, http.StatusNotFound, errors.ErrNotFound)

		return
	}

	if r.Method == "GET" {
		if gr.routes[r.URL.Path].get != nil {
			gr.routes[r.URL.Path].get(w, r)

			return
		}
	}

	if r.Method == "POST" {
		if gr.routes[r.URL.Path].post != nil {
			gr.routes[r.URL.Path].post(w, r)

			return
		}
	}

	errors.JSONError(w, http.StatusMethodNotAllowed,
		fmt.Errorf("%s, %w", r.Method, errors.ErrMethodNotAllowed),
	)
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
