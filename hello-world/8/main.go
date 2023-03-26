package main
//Ponteiros
import "fmt"

func main() {
	//Memoria->Endereço->Valor
	a := 10
	// var ponteiro *int = &a
	// fmt.Println(ponteiro)

	valordex(a) //Soma 1 à variável
	fmt.Println(a) //Imprime a variável, porém a alteração anterior não prevalece
	fmt.Println("O endereço de memória da variável é ", &a)

	ponteirodex(&a) //Soma 1 ao valor alocado no endereço de memória utilizando ponteiro
	fmt.Println(a) // Agora sim a alteração no valor é guardada
}

func valordex(x int) {
	x++
	fmt.Println(x) //Imprime a variável
}

func ponteirodex(x *int) {
	*x++
	fmt.Println("\n",x) //Imprime o Ponteiro
}