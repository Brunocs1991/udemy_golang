package main

import "fmt"

func main() {

	//Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//fim aritimeticos

	//atribuicao
	fmt.Println("--------------")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//fim dos operadores de atribuicao

	//operadores relacionais
	fmt.Println("--------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	//fim dos operadores relacionais

	//operadores logicos
	fmt.Println("--------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//fim operadores logicos

	//operadores unarios
	fmt.Println("--------------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	fmt.Println("--------------")
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)

}
