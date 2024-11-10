package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	Number  string  `json:"number"`
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{Number: "123456", Balance: float64(1000)}

	//Serialize struct as json, save it in a variable and print
	accountJson, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(accountJson))

	//Serialize struct as json and send the value to stdout
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	//Desserialize json to struct, save it in a variable and print
	accountJson = []byte(`{"number":"654321","balance":1500}`)
	err = json.Unmarshal(accountJson, &account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number: %s, Balance: %f\n", account.Number, account.Balance)

	err = json.NewDecoder(strings.NewReader(string(accountJson))).Decode(&account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number: %s, Balance: %f\n", account.Number, account.Balance)
}
