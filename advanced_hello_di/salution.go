package main

import "advanced_hello_di/writers"

type Salution struct {
	writer writers.IMessageWriter
}

func (s Salution) Exclaim() {
	s.writer.Write("Hello, DI!")
}
