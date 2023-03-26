package main
//Generics
import "fmt"

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

//Soma com Generics
func Soma[T int | float64] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

//Constraint
type Number interface{
	~int | ~float64 	//
}
//Generics com constraints
func Soma2[T Number] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}


func main() {
	m := map[string]int{
		"Um": 1,
		"Dois": 2,
		"Tres": 3,
		"Douglas": 4000,
		"João": 6000,
	}

	m2 := map[string]float64{ "Amendoin": 2.5, "Lanche": 23.78, "Refrigerante": 6.5, "Bananada":1.0}

	println(SomaInteiro(m))
	fmt.Println(SomaFloat(m2))
	fmt.Println(Soma(m2))
	fmt.Println(Soma2(m))
}