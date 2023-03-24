package main

import (
	"net/http"
	"text/template"
)

	type Curso struct {
		Nome 			string
		CargaHoraria 	int
	}

	type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html")) //Apenas um arquivo
		//Caso for passar vários arquivos você pode usar um slice do tipo
		//[]string{"template1.html", "template2.html", "template3.html"}
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 70},
			{"C#", 60},
			{"PHP", 37},
			{"Python", 28},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8081", nil)

}