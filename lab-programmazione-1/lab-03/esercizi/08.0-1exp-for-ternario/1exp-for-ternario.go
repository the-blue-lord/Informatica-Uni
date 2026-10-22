package main

import "fmt"

func main() {
	/*
		Definita una costante per la base come 8, chiedi all'utente un numero intero tra 1 e 4 compresi. Stampa quindi un numero di pipes separate da tab pari al prodotto tra l'input e la base in uso. Poi a capo stampa inoltra le conversioni alla base in useo separate da tab dei numeri compresi tra 0  (compreso) e il prodotto tra l'input e la base in uso.
	*/

	const BASE = 8;
	var n int;

	fmt.Print("un int tra 1 e 4: ");
	fmt.Scan(&n);

	for i := 0; i < n*BASE; i++ { // i < n*BASE
		fmt.Print("|\t");
	}
	fmt.Println();

	for i := 0; i < n*BASE; i++ {
		fmt.Print(i/BASE, i%BASE, "\t");
	}
	fmt.Println();
}
