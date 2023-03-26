	package main
//Type Assertion
	import "fmt"

	func main() {
		var interfaceSemTipo interface{} = "Eerling Haaland"
		println(interfaceSemTipo.(string))
		//fmt.Printf("%T\n", interfaceSemTipo)
		res, ok := interfaceSemTipo.(int) //O ok assume tipo como booleano
		fmt.Printf("O valor é %v e o resultado de ok é %v \n", res, ok)
		res2 := interfaceSemTipo.(int) //Como aqui não coloquei o ok vair dar panic, Go é fortemente tipada
		fmt.Printf("O valor é %v e o resultado de ok é %v \n", res2, ok) //Ok deve estar na atribuição da variável
	}