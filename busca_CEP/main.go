package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {		
		/*===============================================================================>
		//req, err := http.Get(url)  //for iterando "url"		
		A url deve seguir o formato da API com o CEP entre "/ws/" e "/json/", por exemplo:
		go run busca_CEP/* http://viacep.com.br/ws/29890-000/json/					
		A requisição recebe uma url por padrão, porém é possível manipulá-la para o usuário digitar apenas o CEP: 
		<=================================================================================*/
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")	//for iterando "cep"
		if err != nil {
			//panic(err)
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err) //Printf aqui não funciona
		}
		defer req.Body.Close()


		res, err := io.ReadAll(req.Body)
		if err != nil {
			//panic(err)
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			//panic(err)
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse(Unmarshal) da resposta: %v\n", err)
		}
		//fmt.Println(data.Localidade)
		file, err := os.Create("cidade.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade: ", data.Localidade)
	}
}