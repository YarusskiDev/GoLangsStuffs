package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`

	//caso eu queria que uma propriedade seja ignorada é só colocar um - no nome da propridade
}

func main() {

	cachorroEmJson := `{"nome":"cristiano","raca":"humano","idade":31}`
	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorroEmJson2 := `{"nome":"BRUNO","gordo":"humano"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJson2), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}
