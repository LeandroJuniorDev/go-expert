package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Address struct {
	ZipCode      string `json:"cep"`
	Address      string `json:"logradouro"`
	Address2     string `json:"complemento"`
	Suite        string `json:"unidade"`
	Neighborhood string `json:"bairro"`
	Location     string `json:"localidade"`
	FU           string `json:"uf"`
	State        string `json:"estado"`
	Region       string `json:"regiao"`
	Ibge         string `json:"ibge"`
	Gia          string `json:"gia"`
	DDD          string `json:"ddd"`
	Siafi        string `json:"siafi"`
}

func main() {

	http.HandleFunc("/address/{zipcode}", getAddress)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getAddress(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.URL.Path,  "/address") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipCode := r.PathValue("zipcode")
	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := findAddress(zipCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set( "Content-Type",  "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func findAddress(zipCode string) (*Address, error) {
	res, err := http.Get( fmt.Sprintf("https://viacep.com.br/ws/%s/json", zipCode))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var address Address
	err = json.Unmarshal(resBody, &address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
