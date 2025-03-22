package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const VIA_CEP_API = "https://viacep.com.br/ws/%s/json/"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/busca-cep", BuscaCEPHandler)
	http.ListenAndServe(":8080", mux)
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP não informado"))
		return
	}
	address, err := BuscaCEP(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar CEP"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(address)
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