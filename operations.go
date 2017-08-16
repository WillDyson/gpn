package main

import "math/cmplx"

type operation struct {
	label       string
	description string
	command     func([]complex128) ([]complex128, error)
}

var (
	add = operation{
		"+",
		"(y+x) adds x and y on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 2 {
				return stack, errStackTooSmall
			}
			x, y := pop(&stack), pop(&stack)
			return append(stack, y+x), nil
		},
	}
	minus = operation{
		"-",
		"(y-x) takes x from y on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 2 {
				return stack, errStackTooSmall
			}
			x, y := pop(&stack), pop(&stack)
			return append(stack, y-x), nil
		},
	}
	multiply = operation{
		"*",
		"(y*x) multiplies x with y on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 2 {
				return stack, errStackTooSmall
			}
			x, y := pop(&stack), pop(&stack)
			return append(stack, y*x), nil
		},
	}
	divide = operation{
		"/",
		"(y/x) divides y by x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 2 {
				return stack, errStackTooSmall
			}
			x, y := pop(&stack), pop(&stack)
			return append(stack, y/x), nil
		},
	}
	power = operation{
		"^",
		"(y**x) raises y to the power of x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 2 {
				return stack, errStackTooSmall
			}
			x, y := pop(&stack), pop(&stack)
			return append(stack, cmplx.Pow(y, x)), nil
		},
	}
	exp = operation{
		"$",
		"(e**x) raises e to the power of x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Exp(x)), nil
		},
	}
	log = operation{
		"log",
		"(log x) logs x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Log(x)), nil
		},
	}
	sin = operation{
		"sin",
		"(sin x) sins x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Sin(x)), nil
		},
	}
	cos = operation{
		"cos",
		"(cos x) coses x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Cos(x)), nil
		},
	}
	tan = operation{
		"tan",
		"(tan x) tans x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Tan(x)), nil
		},
	}
	asin = operation{
		"asin",
		"(asin x) asins x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Asin(x)), nil
		},
	}
	acos = operation{
		"acos",
		"(acos x) acoses x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Acos(x)), nil
		},
	}
	atan = operation{
		"atan",
		"(atan x) atans x on the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) < 1 {
				return stack, errStackTooSmall
			}
			x := pop(&stack)
			return append(stack, cmplx.Atan(x)), nil
		},
	}
	sum = operation{
		"sum",
		"sums all values on the stack",
		func(stack []complex128) ([]complex128, error) {
			var S complex128
			for _, v := range stack {
				S += v
			}
			return []complex128{S}, nil
		},
	}
	mean = operation{
		"mean",
		"takes the mean of all the values on the stack",
		func(stack []complex128) ([]complex128, error) {
			var M complex128
			for _, v := range stack {
				M += v
			}
			return []complex128{M / complex(float64(len(stack)), 0.0)}, nil
		},
	}
	clear = operation{
		"clear",
		"clears the stack",
		func(stack []complex128) ([]complex128, error) {
			return nil, nil
		},
	}
	count = operation{
		"count",
		"counts the number of values on the stack",
		func(stack []complex128) ([]complex128, error) {
			return append(stack, complex(float64(len(stack)), 0.0)), nil
		},
	}
	popop = operation{
		"pop",
		"pops the element at the top of the stack",
		func(stack []complex128) ([]complex128, error) {
			if len(stack) == 0 {
				return stack, errStackTooSmall
			} else {
				return stack[:len(stack)-1], nil
			}
		},
	}
)

var operations_keys = []*operation{
	&add,
	&minus,
	&multiply,
	&divide,
	&power,
	&exp,
	&log,
	&sin,
	&cos,
	&tan,
	&asin,
	&acos,
	&atan,
	&sum,
	&mean,
	&clear,
	&count,
	&popop,
}

var operations = map[string]*operation{
	add.label:      &add,
	minus.label:    &minus,
	multiply.label: &multiply,
	divide.label:   &divide,
	power.label:    &power,
	exp.label:      &exp,
	log.label:      &log,
	sin.label:      &sin,
	cos.label:      &cos,
	tan.label:      &tan,
	asin.label:     &asin,
	acos.label:     &acos,
	atan.label:     &atan,
	sum.label:      &sum,
	mean.label:     &mean,
	clear.label:    &clear,
	count.label:    &count,
	popop.label:    &popop,
}
