package pattern

import (
	"context"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Самое очевидное - middleware в роутерах
*/

// пример доступа к защищенным данным
func chainUsage() {
	token := "user token"
	dataID := "data id"

	// валидация токена
	// проверка прав доступа
	// получение данных из кеша

	// init chain
	var chain Chain
	chain = TokenMiddleware(chain)
	chain = AccessMiddleware(chain)
	chain = CacheMiddelware(chain)

	ctx := context.WithValue(context.Background(), "token", token)
	ctx = context.WithValue(ctx, "dataID", dataID)

	if err := chain(ctx); err != nil {
		log.Println(fmt.Errorf("wrong access: %w", err))
	}
}

type Chain func(ctx context.Context) error

func TokenMiddleware(next func(ctx context.Context) error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		log.Println("token middleware")

		return next(ctx)
	}
}

func AccessMiddleware(next func(ctx context.Context) error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		log.Println("access middleware")

		return next(ctx)
	}
}

func CacheMiddelware(next func(ctx context.Context) error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		log.Println("cache middleware")

		return next(ctx)
	}
}
