// 2020/10/12

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
		"twice":         twice,
		"addSomeNumber": addSomeNumber,
	}

	tpl := template.Must(
		template.New("sample.tpl").Funcs(funcMap).ParseFiles("./sample.tpl"),
	)

	errTplExecute := tpl.Execute(os.Stdout, inputs)

	if errTplExecute != nil {
		panic(errTplExecute)
	}
}

func twice(num int) (int, error) {
	return num * 2, nil
}

func addSomeNumber(x int, y int) (int, error) {
	return x + y, nil
}

// End
