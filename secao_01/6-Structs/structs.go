package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")
	var u usuario
	u.nome = "Bruno"
	u.idade = 2023 - 1991
	fmt.Println(u)

	enderecoExemplo := endereco{"r1", 99}

	usuario2 := usuario{"Patricia", 2023 - 1994, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 99}
	fmt.Println(usuario3)
}
