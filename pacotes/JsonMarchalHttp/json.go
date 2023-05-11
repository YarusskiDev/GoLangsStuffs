package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"cristiano", "humano", 31}
	fmt.Print(c)

	cachorroEmJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

}
