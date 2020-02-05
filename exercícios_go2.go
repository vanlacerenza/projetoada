/*  1:-) Construa um algoritmo que calcule a quantidade de latas de tinta necessárias
 e o custo para pintar cilindricos. 
São fornecidos a altura e o raio desse cilindro.
Sabendo que:
a lata de tinta custa $ 50,00;
cada lata contém 5 litros;
cada litro de tinta pinta 3 metros quadrados.
Dados de entrada: altura (H) e raio (R).
Dados de saída: custo (C) e quantidade (QTDE).
Utilizando o planejamento reverso, sabemos que:

o custo é dado pela quantidade de latas * $ 50,00;
a quantidade de latas é dada pela quantidade total de litros/5;
a quantidade total de litros é dada pela área do cilindro/3;

a área do cilindro é dada pela área da base + área lateral;
a área da base é (PI * R2);
a área lateral é a altura * comprimento: (2 * PI * R * H);
sendo que R (raio) e  H (altura) são dados de entrada e PI é uma constante 
de valor conhecido: 3,1415.
*/

package main

 import "fmt"
  
 func main() {
 
	 var H, R float32 = 5, 2
	 var C, QTDE, PI float32 = 0, 0, 3.1415
 
	 /* ARÉA DO CILINDRO */
 
	 var areaLateral = (2 * PI * R * H)
	 var areaBase = (PI * ((R * R))
	 var areaCilindro = areaBase + areaLateral
	 var litros = areaCilindro / 3
 
	 /* Calculando o que ele pede, por fim
	 não sei usar uma potência, pot não deu boa */
 
	 QTDE = litros / 5
	 C = QTDE * 50
 
	 fmt.Println("a quantidade de latas é", QTDE, "e o custo é", C)
 
 }


 // ssh > add ao manual do git 


 /*2:-) Dados três valores A, B, C, verificar se eles podem ser os comprimentos
  dos lados de um triângulo, se forem, verificar se compõem um triângulo eqüilátero, 
  isósceles ou escaleno. Informar se não compuserem nenhum triângulo.
Dados de entrada: três lados de um suposto triângulo (A, B, C).
Dados de saída – mensagens: não compõem triângulo, triângulo eqüilátero,
 triângulo isósceles, triângulo escaleno.

O que é um triângulo eqüilátero?
Um triângulo com três lados iguais.

O que é um triângulo isósceles?
Um triângulo com dois lados iguais.

O que é um triângulo escaleno?
Um triângulo com todos os lados diferentes. */

package main

import "fmt"

func main() {

	var a, b, c int = 15, 6, 5

	if a == b && a == c {
		fmt.Println("o triângulo é equilatéro")
	} else {
		if a == b || a == c || b == c {
			fmt.Println("o triangulo é isosceles")
		}	
			if a != b && a != c && b != c {
				fmt.Println("o triangulo é escaleno")
			} else {
				if a == 0 || b == 0 || c == 0 {
					fmt.Println("não é um triangulo")
	
				}
			}
		
	}
}



https://golang.org/pkg/math/#example_Pow // potencia