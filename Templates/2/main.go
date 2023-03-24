package main

import (
	"os"
	"text/template"
)

	type Curso struct {
		Nome 			string
		CargaHoraria 	int
	}

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("O curso {{.Nome}} tem carga horária de {{.CargaHoraria}} horas.  \n"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}
}