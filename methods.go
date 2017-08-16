package main

import "fmt"

func top(stack []complex128) complex128 {
	if len(stack) == 0 {
		return 0
	}
	return stack[len(stack)-1]
}

func pop(stack *[]complex128) (r complex128) {
	if len(*stack) == 0 {
		return 0
	}
	r = (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return
}

func convComplex(in string) (complex128, bool) {
	var real, imag float64
	_, err := fmt.Sscanf(in, "%f+i%f", &real, &imag)
	if err == nil {
		return complex(real, imag), true
	}
	_, err = fmt.Sscanf(in, "%f", &real)
	if err == nil {
		return complex(real, 0), true
	}
	_, err = fmt.Sscanf(in, "i%f", &imag)
	if err == nil {
		return complex(0, imag), true
	}
	return 0, false
}
