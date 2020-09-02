package a

type myError struct { // want "OK"
}

func (e myError) Error() string {
	return ""
}

type myError2 struct { // want "OK"
}

func (e *myError2) Error() string {
	return ""
}

type hoge struct {
}

type myError3 struct { // want "OK"
	myError
}
