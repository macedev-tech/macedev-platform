package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	fmt.Println("Argumentos recebidos:")

	for i, arg := range os.Args {
		fmt.Printf("[%d] %s\n", i, arg)
	}

	fmt.Println()

	if len(os.Args) < 2 {
		fmt.Println("Bem-vindo à MaceDev Platform!")
		return
	}

	switch os.Args[1] {

	case "hello":
		fmt.Println("Olá, Valmir! Bem-vindo à MaceDev Platform!")

	default:
		fmt.Println("Comando não reconhecido:", os.Args[1])
	}
}
