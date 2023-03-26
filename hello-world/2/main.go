package main

import "fmt"

type ID int //Criar tipos próprios

var (
	b bool  //false
	c int	//0
	d string = "teste"
	e float64 

	i ID = 777
)
//Tudo que é declarado deve ser usado

func main() {
	//f:= "First attribution"
	println(i)
	fmt.Printf("O tipo da variável é %T \n", i)
}