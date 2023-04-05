package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Ana", 54})
	fmt.Println(pessoa{"José Daniel", 25})

}

//estruturas: coleções de "campos"
