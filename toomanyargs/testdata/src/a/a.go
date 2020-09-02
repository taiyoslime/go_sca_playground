package a

func f1(a, b, c, d int) {}

func f2(a int, b int, c int, d int, e int) {} // want "too many arguments"

func f3(a, b, c, d, e int) {} // want "too many arguments"

type Test struct {}

func (t *Test) f4 (a int, b bool, c int, d bool, e int) interface{}{  // want "too many arguments"
	f5 := func(a, b, c, d, e int) {} // want "too many arguments"
	return f5
}
