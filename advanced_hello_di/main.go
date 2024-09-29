package main

import (
	"advanced_hello_di/writers"
	"os"
	"reflect"
)

var writerConfig = os.Getenv("WRITER")

func main() {
	writerType, found := writers.StructMap[writerConfig]
	if !found {
		panic("writer not found")
	}
	newWriter := reflect.New(writerType).Elem().Interface()
	writer, _ := newWriter.(writers.IMessageWriter)
	salution := Salution{writer}
	salution.Exclaim()
}
