package main

import "hello_di/writers"

func main() {
	writer := writers.ConsoleMessageWriter{}
	salution := Salution{writer}
	salution.Exclaim()
}
