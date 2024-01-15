package main

import (
	"fmt"
	"time"
)

var foo int

func main() {
	incVar(&foo)
}

func incVar(varPtr *int) {
	for {
		*varPtr++
		fmt.Printf("Valor: %d | Endereço: 0x%X\n", *varPtr, varPtr)
		time.Sleep(time.Second)
	}
}
