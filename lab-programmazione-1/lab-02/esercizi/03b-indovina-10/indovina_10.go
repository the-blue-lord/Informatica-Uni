package main

import "fmt"

func main() {
	const NUMERO_SEGRETO = 7;

	var n int;

	fmt.Print("indovina il numero che ho pensato: ");
	fmt.Scan(&n);

	if(n < 1 || n > 10) {
		fmt.Println("Valore non valido");
	} else if(n != NUMERO_SEGRETO) {
		fmt.Println("Non hai indovinato");
	} else {
		fmt.Println("Hai indovinato!");
	}
}