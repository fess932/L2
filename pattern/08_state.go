package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Поведенческий паттерн, неотьемлемо связан с машиной состояний
	поволяющий поведению объекта изменяться в зависимости от его состояния.

плюсы упрощает сложные процессы с множеством вариантов ветвления, добавляет слои

минусы может черезмерно усложнить программу если ветвлений не слишком много
или они меняются не часто
*/

// Есть контекст и контекст имеет некоторое состояние.
// Есть список состояний которые реализуют интерфейс Состояния.
// На примере онлайн кассы

func checkErr(err error) {
	if err != nil {
		log.Println("ошибка:", err)
	}
}

func stateUsage() {
	var kkt IState = NewKKTStateContext()

	checkErr(kkt.Payment(123))
	checkErr(kkt.CloseShift())
	checkErr(kkt.OpenShift())
	checkErr(kkt.Payment(123))
	checkErr(kkt.OpenShift())
	checkErr(kkt.CloseShift())
}

type IState interface {
	OpenShift() error        // открыть смену
	Payment(price int) error // оплата
	CloseShift() error       // закрытие смены
}

type KKTStateContext struct {
	shiftOpened IState
	shiftClosed IState

	currentState IState
}

func (k *KKTStateContext) setState(state IState) {
	k.currentState = state
}

func (k *KKTStateContext) OpenShift() error {
	return k.currentState.OpenShift()
}
func (k *KKTStateContext) Payment(price int) error {
	return k.currentState.Payment(price)
}
func (k *KKTStateContext) CloseShift() error {
	return k.currentState.CloseShift()
}

func NewKKTStateContext() *KKTStateContext {
	k := &KKTStateContext{}

	k.shiftClosed = &KKTClosedState{k}
	k.shiftOpened = &KKTOpenedState{k}
	k.currentState = k.shiftClosed

	return k
}

// список состояний

// Касса закрыта
type KKTClosedState struct {
	sctx *KKTStateContext
}

func (k *KKTClosedState) OpenShift() error {
	k.sctx.setState(k.sctx.shiftOpened)
	fmt.Println("Смена открыта")

	return nil
}
func (k *KKTClosedState) Payment(price int) error {
	return fmt.Errorf("Касса закрыта")
}
func (k *KKTClosedState) CloseShift() error {
	return fmt.Errorf("Касса закрыта")
}

// Касса открыта
type KKTOpenedState struct {
	sctx *KKTStateContext
}

func (k *KKTOpenedState) OpenShift() error {
	return fmt.Errorf("Смена уже открыта")
}
func (k *KKTOpenedState) Payment(price int) error {
	fmt.Println("Оплата прошла успешно", price)
	return nil
}
func (k *KKTOpenedState) CloseShift() error {
	k.sctx.setState(k.sctx.shiftClosed)
	fmt.Println("Смена закрыта")
	return nil
}
