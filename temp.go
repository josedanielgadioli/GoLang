// declarar meu pacote principal
package main

//importar a função fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoF = 212.0

//declarar função principal
func main() {

	F := ebulicaoF        //variável do valor da temperatura em graus F
	C := (F - 32) * 5 / 9 //conversão de F para C
	//apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em °F = %g (%T) e a temperatura de ebulição da água em °C = %g (%T).", F, F, C, C)
	//a gente espera que apareça F = 212 e C = 100

}
