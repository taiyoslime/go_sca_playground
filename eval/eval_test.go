package eval

import (
	"testing"
)

func TestEval(t *testing.T) {
	v, err := Eval("1 - 1 + -2 * 3 / 2")
	if v != -3 || err != nil {
		t.Error()
	}
}
