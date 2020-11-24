// 2020/10/12

package main

import (
	"fmt"
	"os"
	"text/template"
)

type inputValues struct {
	Name   string
	Number int
	String string
}

func main() {
	inputs := inputValues{
		Name:   "world",
		Number: 5,
		String: "Hello",
	}

	funcMap := template.FuncMap{
		"twice":         twice,
		"addSomeNumber": addSomeNumber,
		"joinStrings":   joinStrings,
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

func joinStrings(addString string, base string) (string, error) {
	return fmt.Sprintf("%s %s", base, addString), nil
}

// End
