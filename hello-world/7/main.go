package main

//Struct
import "fmt"

type Address struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	DesativarCliente()
	AtivarCliente()
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	//Address
}

func (c Client) DesativarCliente() {
	c.Ativo = false
	fmt.Printf("Cliente %s desativado \n Status: %t \n", c.Nome, c.Ativo)
	//Não está salvando na struct
}

func Desativacao(pessoa Pessoa) {
	pessoa.DesativarCliente()
}

func (c Client) AtivarCliente() {
	c.Ativo = true
	fmt.Printf("Cliente %s ativado \n Status: %t \n", c.Nome, c.Ativo)
	//Não está salvando na struct
}

func main() {
	cliente1 := Client{"Wesley Safadão", 31, false}
	cliente2 := Client{Nome: "Wesley", Idade: 21, Ativo: false}
	pr9 := Client{
		Nome:  "Pedro Raul",
		Idade: 25,
		Ativo: true,
	}

	cliente1.AtivarCliente()

	pr9.DesativarCliente()
	Desativacao(pr9) 

	fmt.Println(cliente1)
	fmt.Println(cliente2, pr9)

	fmt.Println("Nome do cliente 1 é", cliente1.Nome) //funcionou sem o "%s" mas o certo é colocar
}
