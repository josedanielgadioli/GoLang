package main

import "fmt"

func main() {
	//declarar variável, limite de range e comportamente na função
	for i := 1; i <= 100; i++ {
		//definir o laço
		if i%3 == 0 {
			fmt.Println(i, "é divisível por 3.")
		}
	}
}

//espera-se que imprima (i) é divisível por 3.
