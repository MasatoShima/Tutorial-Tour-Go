//2020/10/12

package main

import (
	"os"
	"text/template"
)

type inputValues struct {
	Name string
}

func main() {
	tpl := template.Must(template.ParseFiles("./sample.tpl"))

	inputs := inputValues{
		Name: "world",
	}

	errTplExecute := tpl.Execute(os.Stdout, inputs)

	if errTplExecute != nil {
		panic(errTplExecute)
	}
}

// End
