package main

import "fmt"

func main() {
	/*
		Scrivi un programma che legge un voto intero e se il valore non appartiene all'intervallo chiuso [0, 30], stampa "voto non valido"
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
