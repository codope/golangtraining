package main

import (
	"fmt"
	"os"
	"text/template"
)

type BookInfo struct {
	Title  string
	Author string
}

func main() {
	const tmpl = `
{{range .}}
	{{.Title}} : {{.Author}}
{{end}}
`

	var myuserlist = []BookInfo{
		{"Learning Go", "Sagar Sumit"},
		{"Fun with Go", "Sampathu M"},
		{"Hello Gopher", "Prashanth Surya"},
	}

	t := template.Must(template.New("tmpl").Parse(tmpl))

	err := t.Execute(os.Stdout, myuserlist)
	fmt.Println(err)
}
