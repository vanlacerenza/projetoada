 /*EX. STRING:
 
 func main() {
    var nome string = "Douglas"
    var idade int = 24
	var versao float32 = 1.1*
	
	Para saber se o Go conseguir inferir corretamente o tipo das variáveis,
	 podemos descobri-los, importando o pacote reflect e chamando a sua função TypeOf
	 , passando para ela a variável que queremos saber o tipo:

package main

import "fmt"
import "reflect"

func main() {
    var nome = "Douglas"
    var idade = 24
    var versao = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}
	Para deixar o nosso código mais limpo ainda, 
	podemos remover a palavra var das variáveis. Podemos fazer isso pois o 
	Go possui um segundo operador de atribuição de variáveis, 
	um mais "curto", que é o :=.


func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
	/
	