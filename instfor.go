package main

import (
	"fmt"
)

func main() {

	for x := 0; x < 20; x++ {
		if x == 3 {
			continue
		} else {
			fmt.Println(x)
		}
	}

}

//SE o número é PAR ou ÍMPAR
//em "GO", o resto é representado por %
