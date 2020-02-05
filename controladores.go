package main

import(
	"fmt"
)

// importa um pacote de funções

// https://play.golang.org/ para praticar

// https://gobyexample.com/ 

func main() {
	fmt.Println("Hello") // ln > pula linha, fmt> é o pacote de funções.


	/* Escreva um algoritmo que armazene o valor 10 em uma variável A 
	e o valor 20 em uma variável B. 
	A seguir (utilizando apenas atribuições entre variáveis) 
	troque os seus conteúdos fazendo com que o valor que está em A 
	passe para B e vice-versa. Ao final, escrever os valores que ficaram 
	armazenados nas variáveis.  */


	func main() {
		var a, b int = 10,
			20;
		var trocar int;

		trocar = a
		a = b
		b = trocar
		fmt.Println(a, b)

	}

	/*  2 valores inteiros, quem é o maior */


	func main() {

		var x, y int = 10,
			20
	
		if x > y {
			fmt.Println("x maior que y")
		}
	
		if y > x {
			fmt.Println("y é > que x")
		}
		if y == x {
			fmt.Println("y e x são iguais")
		}
	}