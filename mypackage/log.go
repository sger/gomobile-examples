package mypackage

import "fmt"

func Log(tag string, msg string) string {
	return "[" + tag + "] " + msg
}

func LogError(tag string, str string) (string, error) {
	return Log(tag, str), nil
}

func TestCallback(name string) {
	callback.CallFromGo("Hello " + name + " from Go")
}

var callback Callback

type Callback interface {
	CallFromGo(string)
}

func RegisterCallback(c Callback) {
	callback = c
	fmt.Println("Callback registered", callback)
}
