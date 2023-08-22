package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	eval "github.com/PieerePi/gople/ch7/e7.13-e7.14"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("type in an expression: ")
		expression, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "read expression error, %v", err)
			}
			fmt.Println()
			return
		}
		expression = expression[:len(expression)-1]
		expr, err := eval.Parse(expression)
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse expression error, %v\n", err)
			return
		}

		fmt.Printf("type in variables (<var>=<val>, eg: x=2 y=3): ")
		variables, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "read variables error, %v", err)
			}
			fmt.Println()
			return
		}
		variables = variables[:len(variables)-1]
		env := eval.Env{}
		assignments := strings.Fields(variables)
		for _, a := range assignments {
			fields := strings.Split(a, "=")
			if len(fields) != 2 {
				fmt.Fprintf(os.Stderr, "bad assignment: %s, ignore it\n", a)
				continue
			}
			ident, valStr := fields[0], fields[1]
			val, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "bad value for %s, using zero: %v\n", ident, err)
				continue
			}
			env[eval.Var(ident)] = val
		}

		fmt.Printf("result is: %g\n", expr.Eval(env))
	}
}
