package main

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
}

const tp = `
{{range .}}
	{{if .Edad}}
	Nombre: {{.Nombre}} Edad: {{.Edad}} - Correcto
	{{else}}
	Nombre: {{.Nombre}} Edad: {{.Edad}} - Falso
	{{end}}
{{end}}
`

func main() {

	personas := []Persona{
		{
			Nombre: "Rodrigo",
			Edad:   28,
		},
		{
			Nombre: "Pedro",
			Edad:   12,
		},
		{
			Nombre: "Juan",
			Edad:   0,
		},
	}

	t := template.New("personas")
	t.Parse(tp)
	t, err := t.Parse(tp)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, personas)
	if err != nil {
		panic(err)
	}

}
