// primeiramente declaramos o pacote principal
package main

import "fmt"

//depois importamos a função fmt

//declarar a variável CONST do ponto de ebulição em Kelvin
const ebulicaoK = 372.2

//função principal
func main() {

	Kelvin := ebulicaoK       //variável do valor de temperatura em graus Kelvin
	Celsius := (Kelvin - 273) //conversão de Kelvin para Celsius
	nome := "José Daniel"

	//apareça na tela o meu nome e o que mostrarei
	fmt.Printf("Meu nome é %s e abaixo descreverei o ponto de ebulição da água em graus Celsius e Kelvin", nome)

	//faz o resultado aparecer na tela
	fmt.Printf("\nA temperatura de ebulição da água em °K é = %g K e em °C é %f C.", Kelvin, Celsius)

	//o resultado esperado é Kelvin = 372,2 e C 99,8

}
