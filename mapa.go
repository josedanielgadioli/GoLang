package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])

}

//mapa: coleção ordenada (lista) pares chave-valor
//var x map[string]int
//traduzindo: x é um mapa de strings para ints
