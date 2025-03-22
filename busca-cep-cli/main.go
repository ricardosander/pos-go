package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const VIA_CEP_API = "https://viacep.com.br/ws/%s/json/"

func main() {
	file, err := os.Create("ceps.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		panic(err)
	}
	defer file.Close()

	for _, cep := range os.Args[1:] {
		address, err := BuscaCEP(cep)
		if err != nil {
			fmt.Println("Erro ao buscar CEP:", err)
			continue
		}
		file.WriteString(address.ToString() + "\n")
	}
}

func BuscaCEP(cepInput string) (*ViaCEPResponse, error) {
	resp, err := http.Get(fmt.Sprintf(VIA_CEP_API, cepInput))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err :=  io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cep ViaCEPResponse
	json.Unmarshal(data, &cep)
	return &cep, nil
}

type ViaCEPResponse struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro string `json:"bairro"`
	Cidade string `json:"localidade"`
	Estado string `json:"uf"`
}

func (v ViaCEPResponse) ToString() string {
	return fmt.Sprintf("%s, %s - %s, %s", v.Logradouro, v.Bairro, v.Cidade, v.Estado)
}