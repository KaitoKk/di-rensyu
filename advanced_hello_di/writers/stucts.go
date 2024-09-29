package writers

import "reflect"

var StructMap = map[string]reflect.Type {
	"ConsoleMessageWriter": reflect.TypeOf(ConsoleMessageWriter{}),
}
