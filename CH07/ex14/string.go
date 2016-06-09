// Copyright © 2016 "Shun Yokota" All rights reserved

package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprint(float64(l))
}

func (u unary) String() string {
	if isUnaryOrBinary(u.x) {
		return string(u.op) + stringWithPrimary(u.x)
	}
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	var x string
	var y string
	if isUnaryOrBinary(b.x) {
		x = stringWithPrimary(b.x)
	} else {
		x = b.x.String()
	}
	if isUnaryOrBinary(b.y) {
		y = stringWithPrimary(b.y)
	} else {
		y = b.y.String()
	}
	//fmt.Printf("%s %s %s\n", x, string(b.op), y)
	return x + " " + string(b.op) + " " + y
}

func (c call) String() string {
	var args string
	for n, ex := range c.args {
		if isUnaryOrBinary(ex) && len(c.args) != 1 {
			args += stringWithPrimary(ex)
		} else {
			args += ex.String()
		}
		if n < len(c.args)-1 {
			args += ", "
		}
	}
	return c.fn + "(" + args + ")"
}

func (mn min) String() string {
	return call{"min", mn}.String()
}

func stringWithPrimary(x Expr) string {
	return "(" + x.String() + ")"
}

func isUnaryOrBinary(x Expr) bool {
	_, isUnary := x.(unary)
	_, isBinary := x.(binary)
	return isUnary || isBinary
}
