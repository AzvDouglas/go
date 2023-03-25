package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Adicione esta linha para imprimir as mensagens de log no objeto ResponseWriter (Terminal do client)
	//log.SetOutput(w) //Print para o client
	log.Println("Request iniciada")
	defer log.Println("Request Finalizada")
	select {
	case <-time.After(7 * time.Second):
		log.Println("Request processada com sucesso!")
		w.Write([]byte("Request printada para o Client \n"))
	case <-ctx.Done():
		// err := ctx.Err()
		log.Println("Request cancelada pelo cliente:")
	}
}