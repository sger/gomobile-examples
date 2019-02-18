package mypackage

func Log(tag string, msg string) string {
	return "[" + tag + "] " + msg
}

func LogError(tag string, str string) (string, error) {
	return Log(tag, str), nil
}

var callback Callback

type Callback interface {
	Method1()
	Method2(string)
}

func RegisterCallback(c Callback) {
	callback = c
}
