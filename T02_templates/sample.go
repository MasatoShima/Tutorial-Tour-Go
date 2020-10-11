//2020/10/12

package main

import (
	"os"
	"text/template"
)

type inputValues struct {
	Name   string
	Number int
}

func main() {
	inputs := inputValues{
		Name:   "world",
		Number: 5,
	}

	funcMap := template.FuncMap{
		"twice":        twice,
		"addAnyNumber": addSomeNumber,
	}

	tpl := template.Must(
		template.New("sample.tpl").Funcs(funcMap).ParseFiles("./sample.tpl"),
	)

	errTplExecute := tpl.Execute(os.Stdout, inputs)

	if errTplExecute != nil {
		panic(errTplExecute)
	}
}

func twice(num int) int {
	return num * 2
}

func addSomeNumber(x int, y int) int {
	return x + y
}

// End
