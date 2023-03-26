package main
//Slices
import "fmt"

func main() {
	s := []int{2, 7, 11, 29, 37}
	fmt.Printf("O tipo da variável é %T \n", s)
	fmt.Printf("O último item do slice é %d \n", s[len(s) - 1])
	fmt.Printf("O tamanho do slice é %d \n", len(s))
	fmt.Printf("A capacidade do slice é %d \n", cap(s))

	s = append(s, 42)
	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])
}