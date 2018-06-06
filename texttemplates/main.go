package main

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
	Tipo   string
}

/*const tp = `
{{range .}}
	{{if .Edad}}
	Nombre: {{.Nombre}} Edad: {{.Edad}} - Correcto
	{{else}}
	Nombre: {{.Nombre}} Edad: {{.Edad}} - Falso
	{{end}}
{{end}}
`*/

func main() {

	/*personas := []Persona{
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
	}*/

	persona := Persona{
		Nombre: "Rodri",
		Edad:   28,
		Tipo:   "admin",
	}

	t := template.New("persona")
	t, err := t.ParseGlob("templates/*.txt") // Tambien se puede utilizar Parse File
	if err != nil {
		panic(err)
	}

	/*t.Parse(tp)
	t, err := t.Parse(tp)
	if err != nil {
		panic(err)
	}*/

	/*err = t.Execute(os.Stdout, personas)
	if err != nil {
		panic(err)
	}*/

	err = t.ExecuteTemplate(os.Stdout, "usuario", persona)
	if err != nil {
		panic(err)
	}

}
