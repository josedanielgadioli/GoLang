//método é uma função associada a um tipo particular
//isto é, em POO(Prog. Orientada a Objeto), objeto e um valor(variável) e o método é uma função associada a esse objeto

package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

//este método 'área' possui um tipo 'retângulo
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

//Métodos podem ser definidos por qualquer tipo de receptor
//ponteiro ou valor. Aqui um exemplo do receptor de valor.
func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 12, altura: 3}

	//aqui chamamos os 2 métodos definidos para a nossa estrutura.
	fmt.Println("Área:", r.area())
	fmt.Println("Perímetro:", r.perimetro())
}
