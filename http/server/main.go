package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte("O parâmetro 'cep' é obrigatório e deve conter 8 dígitos"))
		return
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError) //500
		w.Write([]byte("Ocorreu um erro ao buscar o CEP informado"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Hello, World!\n"))
	/*==============================>> Salvando o resultado numa variável <<======================
	result, err := json.Marshal(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //500
		return
	}
	w.Write(result)
	===============================================================================================*/
	json.NewEncoder(w).Encode(cep)
}

//Se o tipo de retorno não for um ponteiro não será possível retornar nil
//cannot use nil as ViaCEP value in return statement
func BuscaCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error	
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var data ViaCEP
	error = json.Unmarshal(body, &data)
	if error != nil {
		return nil, error
	}
	//Retorna a localização da variável na memória 
	return &data, nil
}

//go run http/server/main.go
//curl  localhost:8080
//curl  localhost:8080?cep=29890000
//http://localhost:8080/?cep=29890000
