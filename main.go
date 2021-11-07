package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Dado struct {
	Name        string
	LastName    string
	PhoneNumber string
	Age         string
}

var Dados [5]Dado

var i = 0
var b = 1

func main() {

	//Abrindo o ficheiro
	f, err := os.Open("dados.txt")

	if err != nil {
		log.Fatal(err)
	}

	r, err := io.ReadAll(f)

	for _, value := range r {
		letter := string(value)
		if letter == "\n" {
			break
		}
		if letter == " " {
			b++
		}
		if b == 2 {
			Dados[i].LastName += letter

		}
		if b == 3 {
			Dados[i].PhoneNumber += letter

		}
		if b == 4 {
			Dados[i].Age += letter

		}
		if b == 5 {
			b = 1
			i++
		}
		if letter != "" && b == 1 {
			Dados[i].Name += letter
		}

	}

	f.Close()

	dadosJson, err := json.Marshal(Dados[0:4])
	if err != nil {
		log.Fatal(err.Error())
	}
	takeJson := string(dadosJson)
	fmt.Println(takeJson)

	//abrir ficheiro para escrita
	file, err := os.OpenFile("db.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// escrever no ficheiro
	conteudo := takeJson
	escreve, err := file.Write([]byte(conteudo))

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("os bytes escritos foram: %d ", escreve)

}
