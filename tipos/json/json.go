package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:""nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1899.90,
		Tags: []string{
			"Promoção",
			"Eletrônico",
		},
	}

	p1Json, _ := json.Marshal(p1) // json, error => valores retornados
	fmt.Println(string(p1Json))

	//json to struct

	var p2 produto
	jsonString := `{"id":2, "nome": "Caneta", "preco": 8.90, "tags":["Papelaria", "Importando"]}`

	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
