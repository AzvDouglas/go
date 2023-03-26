package main

//Funções
import (
	"errors"
	"fmt"
)

func main() {
		
	valor, err := somar(400, 1600)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
		fmt.Println("Salario maior que o anterior")
	}

	fmt.Println(soma_total(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func somar(a, b int) (int, error) {
	if a + b >= 2400 {
		return a + b, nil
	} 
	return a + b, errors.New("Salario menor que o anterior")

}

//Função variádica
func soma_total(valores ...int) int {
	total := 0
	for _, valor := range valores {
		total += valor
	}
	return total
}