package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	eval "github.com/PieerePi/gople/ch7/e7.13-e7.14"
)

func parseEnv(s string) eval.Env {
	env := eval.Env{}
	assignments := strings.Fields(s)
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
	return env
}

func main() {
	// curl -k localhost:8080 -d "expr=avg[1,x]&env=x=2 y=3"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.FormValue("expr"))
		// fmt.Println(r.FormValue("env"))
		exprStr := r.FormValue("expr")
		if exprStr == "" {
			http.Error(w, "no expression", http.StatusBadRequest)
			return
		}
		expr, err := eval.Parse(exprStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		env := parseEnv(r.FormValue("env"))
		fmt.Fprintln(w, expr.Eval(env))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
