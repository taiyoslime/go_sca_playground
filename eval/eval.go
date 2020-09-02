package eval

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func insp(expr ast.Expr) (float64, error) {
	switch n := expr.(type) {
	case *ast.ParenExpr:
		return insp(n.X)
	case *ast.BasicLit:
		val, err := strconv.ParseFloat(n.Value, 64)
		if err != nil {
			return 0, err
		}
		return val, nil
	case *ast.BinaryExpr:
		l, err := insp(n.X)
		if err != nil {
			return 0, err
		}
		r, err := insp(n.Y)
		if err != nil {
			return 0, err
		}
		switch n.Op {
		case token.ADD:
			return l + r, nil
		case token.SUB:
			return l - r, nil
		case token.MUL:
			return l * r, nil
		case token.QUO:
			return l / r, nil
		default:
			return 0, errors.New("unknown operator")
		}
	case *ast.UnaryExpr:
		switch n.Op {
		case token.SUB:
			i, err := insp(n.X)
			if err != nil {
				return 0, err
			}
			return -i, nil
		}

	default:
	}
	return 0, errors.New("unknown expr")
}

func Eval(str string) (float64, error) {
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return 0, err
	}
	result, err := insp(expr)
	if err != nil {
		return 0, err
	}
	return result, nil
}
