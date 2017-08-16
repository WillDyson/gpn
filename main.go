package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	var stack []complex128
	var input string
	var err error

	for {
		fmt.Scanf("%s", &input)
		if input == "." {
			if len(stack) == 0 {
				fmt.Println("(*.******)")
			} else {
				z := top(stack)
				if imag(z) == 0.0 {
					fmt.Printf("(%f)\n", real(z))
				} else {
					fmt.Printf("(%f+i%f)\n", real(z), imag(z))
				}
			}
		} else if input == ".q" {
			return
		} else if input == ".com" || input == "?" {
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
			for _, v := range operations_keys {
				fmt.Fprintln(w, v.label, "\t", v.description)
			}
			fmt.Fprintln(w, ". \t prints the top value of stack")
			fmt.Fprintln(w, ".q \t quits the program")
			fmt.Fprintln(w, ".com \t lists all availiable commands")
			w.Flush()
		} else if c, ok := convComplex(input); ok {
			stack = append(stack, c)
		} else if op, ok := operations[input]; ok {
			stack, err = op.command(stack)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Oops - unrecognized operation (.com for availible operations)")
		}
	}
}
