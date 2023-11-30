package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INt32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = utin8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 11111111111111111111111111111111111111111111111111111111111111111111111111.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'b'
	fmt.Println(char)

	var texto int64
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
