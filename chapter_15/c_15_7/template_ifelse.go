package main

import (
	"os"
	"text/template"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Whill not print. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}}Will print. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("If-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}
