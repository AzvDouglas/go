package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int 		`json:"n"`
	Saldo  float64	`json:"s"`

}
func main() {
	
	conta := Conta{Numero: 01, Saldo: 100}
	res, err := json.Marshal(conta) //Guarda o valor
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) //Já retorna o valor sem salvar
	if err != nil {
		panic(err)
	}
	
	jsonPuro := []byte(`{"n":2,"s":2400}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaX.Saldo)
}