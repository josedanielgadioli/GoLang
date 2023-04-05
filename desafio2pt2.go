package main

import "fmt"

func main() {
	//declarar variável, limite de range e comportamente na função
	for i := 0; i <= 100; i++ {
		//definir o laço
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}

//espera-se que quando i%3 == 0 e i%5 == 0 > PinPan
//espera-se que quando i%3 == 0 > Pin
//espera-se que quando i%5 == 0 > Pan
