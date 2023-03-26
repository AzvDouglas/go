package main

import "fmt"

//Maps
func main() {

	salarios := map[string]int{
		"João":  1500,
		"Maria": 2500,
		"Pedro": 3000,
		"Wesley": 10000,
	}

	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["wesley"] = 5000
	fmt.Println(salarios["wesley"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salário do %s é %d \n", nome, salario) 
	}

	//Blank identyfier
	for _, salario := range salarios {
		fmt.Printf("O salário é %d \n", salario)
	}
}