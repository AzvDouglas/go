package main
//Arrays
import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 33

	fmt.Printf("O tipo da variável é %T \n", meuArray)
	fmt.Println("O último item do array é ", meuArray[len(meuArray) - 1])

	for i, v := range meuArray {
		fmt.Printf("O valor %d está na posição %d \n", v, i)
	}
}