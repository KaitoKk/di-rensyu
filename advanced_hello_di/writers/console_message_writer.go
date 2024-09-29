package writers

import "fmt"

type IMessageWriter interface {
	Write(text string)
}

type ConsoleMessageWriter struct {
}

func (cmw ConsoleMessageWriter) Write(text string) {
	fmt.Println(text)
}
