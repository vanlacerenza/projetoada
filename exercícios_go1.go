/* exercícios de cálculo/variáveis/constantes,inteiros e floats*/

/* 1:-) Faça um algoritmo que receba três notas de um aluno, 
calcule a média aritmética entre as três notas e exiba mensagem de ’Aprovado’ 
ou ’Reprovado’,
considerando a média de aprovação maior ou igual a 7,0. */

package main

import "fmt"

func main() {

    var a, b, c int = 10, 6, 5

    var média int = (a + b + c) / 3
	fmt.Println(média)
	
	if (média >= 7) {
		fmt.Println("Aprovado");
	} else {
		if (média < 7) {
			fmt.Println("Reprovado");
   
}
}
}

/*2:-) Dados dois números inteiros e distintos,
 construa um algoritmo que seja capaz de definir qual é o maior elemento. */

 package main

 import "fmt"
 
 func main() {
 
		 var x, y int = 10, 20
	 
		 if x > y {
			 fmt.Println("x maior que y")
		 }
	 
		 if y > x {
			 fmt.Println("y é > que x")
		 }
		 if y == x {
  }
 }

 
/* 3:-) Após o conhecimento de um numerador e um denominador,
 construa um algoritmo para fornecer o resto desta divisão. */


 package main

 import "fmt"
 
  func main() {
  
		  var x, y int = 10, 20
		  var resto int = y%x
		  fmt.Println(resto)
   }
 
   /* 4:-) Dados dois números inteiros, construa um algoritmo que seja capaz 
   de definir se estes são iguais, e caso isso não ocorrer, qual é o menor elemento.*/ 

   package main

   import "fmt"
   
   func main() {
  
	  var a, b int = 10, 20
  
	  if a == b {
		  fmt.Println("os elementos são iguais")
	  }
  
	  if a < b {
		  fmt.Println("a é < que b")
	  }
	  if b < a {
		  fmt.Println("b é menor que a")
	  }
  }

 /* 5:-) Construa um algoritmo que seja capaz de definir 
 qual o maior elemento entre três conhecidos. */

 package main

 import "fmt"
 
 func main() {

	var a, b, c int = 10, 20, 30

	if (a > b && a > c) {
		fmt.Println("a é > que b e c ")
	} 

	if (b > a && b > c) {
			fmt.Println("b é > que a e c")
	}

	if (c > a && c > b) {
		fmt.Println("c é > que a e b")
		
	}
 }

 /*6:-) Elaborar um algoritmo, que faça a conversão de graus Fahrenheit para Celsius. O algoritmo deve
ler um valor em graus Fahrenheit e transformá-lo através da fórmula:
C = (F-32)*5/9 */

package main

 import "fmt"
 
 func main() {

	var F int = 20 
	var C int = (F-32)*5/9
	fmt.Println(C)

	}
 
	//A Seguir, declarar Float 

	/*7:-) Faça um algoritmo que receba duas notas de um aluno e 
	seus respectivos pesos,
	calcule e exiba a média ponderada dessas notas.*/

package main

import "fmt"

func main() {

    var nota1, nota2 float32 = 1.1 , 2.2
    var média float32 = (nota1 + nota2) / 2
    fmt.Println(média)
}

/*
8:-) Faça um algoritmo que receba o valor de um depósito e o valor da taxa de juros. 
Calcule e exiba o valor do depósito, 
o valor do rendimento e o valor total depois do rendimento.  
J = C . i . t 
*/

package main

import "fmt"

func main() {

	var Capital, taxaDeJuros float32 = 100, 0.1
	var valorFinal float32 = Capital * taxaDeJuros 
   	fmt.Println(valorFinal)
}


/*9:-) Faça um algoritmo que receba a idade de uma pessoa em anos,
 calcule e exiba essa idade em meses, dias, horas e minutos. */


package main

import "fmt"

func main() {

	var idade int = 29
	var meses int = 12 * idade
	var dias int = 365 * idade
	var horas int = 8760 * idade
	var minutos int = 525600 * idade
	   fmt.Println(idade)
	   fmt.Println(meses)
	   fmt.Println(dias)
	   fmt.Println(horas)
	   fmt.Println(minutos)
	   
}

/*
10:-) Faça um algoritmo que receba o salário de um funcionário, 
calcule e exiba o valor do imposto de renda a ser pago, 
sabendo que o imposto equivale a 5% do salário, e também, 
o valor líquido do salário a receber.  */

package main

import "fmt"

func main() {

	var salario, taxaImposto float32 = 1000, 0.05
	var impostoTotal float32 = salario * taxaImposto
	var salarioFinal float32 = salario - impostoTotal
	fmt.Println(salarioFinal)
	   
}

/*  C O N T A C T E N A Ç Ã O

func main() {
    var nome string = "Douglas"
    var idade int
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}
A variável idade está vazia propositalmente, e ao executar o programa, temos a seguinte saída:

resultado:

Olá sr. Douglas sua idade é 0
Este programa está na versão 1.1

*/

/*11:-) Faça um algoritmo que receba o salário de um funcionário,
 calcule e imprima o novo salário sabendo-se que este sofreu um aumento de 25%,
  e também, o salário anterior e o reajuste, separadamente. */

  package main

  import "fmt"
  
  func main() {
  
	  var salario, PorcentagemAumento float32 = 1000, 0.25
	  var aumento float32 = salario * PorcentagemAumento
	  var salarioFinal float32 = salario + aumento
	  fmt.Println(salario, salarioFinal)

  }

  /*  Sabe-se que o quilowatt de energia custa um quinto do salário mínimo. 
  Faça um algoritmo que
receba o valor do salário mínimo e a quantidade de quilowatts
 gasta por uma residência. Calcule e
exiba:
 O valor, em reais, de cada quilowatt;
 O valor, em reais, a ser pago por essa residência;
 O novo valor a ser pago por essa residência, a partir de um desconto de 15%.*/

package main

import "fmt"

func main() {

	var valorQuilowatt, salarioMínimo, energiaGasta float32 = 0.2, 1000, 3
	var quilowattsReais float32 = salarioMínimo * valorQuilowatt
	var contaAPagar float32 = salarioMínimo * valorQuilowatt * energiaGasta
		fmt.Println(quilowattsReais, contaAPagar)

	var porcentagemDesconto float32= 0.15
	var desconto float32 = contaAPagar * porcentagemDesconto
	var novoValorAPagar float32 = contaAPagar - desconto
		fmt.Println(novoValorAPagar)

}

/*13:-) Faça um algoritmo que receba o peso de uma pessoa, um valor inteiro, calcule e exiba:
 O peso dessa pessoa em gramas;
 Se essa pessoa engordar 5%, qual será seu novo peso em gramas. */


package main

import "fmt"

func main() {

	var pesoKg int = 60
	var PesoAMais int = 0.05 * 60
	var pesoGramas int = pesoKg * 1000
	var novoPeso int = pesoKg + PesoAMais 
	fmt.Println("O peso em gramas é", pesoGramas)
	fmt.Println("O novo peso é", novoPeso)

}

/*

14:-) Faça um algoritmo que receba o ano de nascimento de uma pessoa e o ano atual. Calcule e
exiba:
 A idade dessa pessoa;
 Essa idade convertida em semanas.

*/

package main

import "fmt"

func main() {

	var anoNascimento, anoAtual int = 1990, 2020
	var Idade int = anoAtual - anoNascimento
	fmt.Println("a idade dessa velha é", Idade)
	var semanasAno int = 365/7
	var IdadeEmSemanas = Idade * semanasAno 
	fmt.Println("a idade dessa pessoa em semanas é", IdadeEmSemanas)

}

/* 15:-) Um comerciante deseja saber qual é o lucro percentual 
que ele está tendo com a venda de
mercadorias.
 Calcule o lucro percentual de uma mercadoria ao serem fornecidos o preço de compra
e o preço de venda da mesma. */

package main

import "fmt"

func main() {

	var PreçoDeCompra, PreçoDeVenda float32 = 20, 32
	var PorcentagemPreçoVenda = PreçoDeVenda * 100 / PreçoDeCompra	
	var PorcentagemLucro = PorcentagemPreçoVenda - 100
	fmt.Println("a porcentadem de lucro é", PorcentagemLucro, "%")

}


/*16:-) Você faz uma aplicação de P reais a uma taxa de juros i constante por um 
período de n meses.
Qual será o montante M após o término da aplicação?

M = P(1 + i) * n

Calcule também qual o lucro em reais e em percentual resultante da aplicação, sabendo-se
que serão descontados 8% de IOF sobre o lucro bruto. */


package main

import "fmt"

func main() {

	var P, i, n float32 = 2000, 1.5, 10
	var M float32 = P * (1+ i) * n 
	fmt.Println(" o montante será" , M)
	var descontoIOF = M * 0.08
	var lucroReais = (M - P) - descontoIOF 
	fmt.Println(" o lucro em reais será" , lucroReais)
	var PercentualResultanteAplicação float32 = 100 - 0.08	
	fmt.Println(" o percentual resultante da aplicação será", PercentualResultanteAplicação )

}


/*17:-) Um trabalhador recebe R$ 20,00 por hora normal trabalhada
 e um acréscimo de 50% sobre este valor, para cada hora extra trabalhada. 
 Calcule qual o valor de seu salário após uma semana de trabalho,
  com jornada de 44 horas semanais. O algoritmo deverá, inicialmente, 
  requisitar ao usuário qual foi a quantidade de horas que o trabalhador 
  realmente trabalhou na semana. */

  package main

  import "fmt"
  
  func main() {
  
	  var horasExtrasRealizadas float32 = 3
	  var hora float32 = 20
	  var PorcentagemAcrescimento float32 = hora * 0.05
	  var horaExtra float32 = hora + PorcentagemAcrescimento
	  var salario float32 = (44 * hora) + (horasExtrasRealizadas * horaExtra)
	  fmt.Println(" o salário será", salario)
  
  }