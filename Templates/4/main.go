package main

import (
	"os"
	"html/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"footer.html",
		"content.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 70},
		{"C#", 60},
		{"PHP", 37},
		{"Python", 28},
	})
	if err != nil {
		panic(err)
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// })
	// http.ListenAndServe(":8081", nil)

}
