package pattern

import (
	"io"
	"os"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


построение обьекта с помощью вспомогательных функций
в го часто применяется, когда нужно построить обьект по разным параметрам
например логгеры
*/

func BuilderUsage() {
	log := NewBuilder().
		SetLevel(1).
		SetName("my logger").
		SetOutput(os.Stdout).
		SetFormatter(func() {}).
		Logger()
	log.Log("Hello world")
}

type Logger struct {
}

func (l *Logger) Log(str string) {}

type LoggerBuilder struct {
	l *Logger
}

func (l *LoggerBuilder) Logger() *Logger {
	return l.l
}

func NewBuilder() *LoggerBuilder {
	return &LoggerBuilder{}
}

func (l *LoggerBuilder) SetLevel(level int) *LoggerBuilder {
	return l
}
func (l *LoggerBuilder) SetName(name string) *LoggerBuilder {
	return l
}
func (l *LoggerBuilder) SetOutput(output io.Writer) *LoggerBuilder {
	return l
}
func (l *LoggerBuilder) SetFormatter(func()) *LoggerBuilder {
	return l
}
