package main

import "fmt"

func main() {
	const VALORE_BENZINA = 0;
	const VALORE_DIESEL = 1;
	const VALORE_ALCOL = 2;
	const VALORE_CHEROSENE = 3;

	const PREZZO_BENZINA = 1.78;
	const PREZZO_DIESEL = 1.98;
	const PREZZO_ALCOL = 1.2;
	const PREZZO_CHEROSENE = 1.1;

	var carburante int;
	var litri float64;

	fmt.Print("tipo di carburante? (0 (benzina), 1 (diesel), 2 (alcol), 3 (cherosene)): ");
	fmt.Scan(&carburante);

	var prezzo_carburante float64;
	if(carburante == VALORE_BENZINA) {
		prezzo_carburante = PREZZO_BENZINA;
	} else if (carburante == VALORE_DIESEL) {
		prezzo_carburante = PREZZO_DIESEL;
	} else if (carburante == VALORE_ALCOL) {
		prezzo_carburante = PREZZO_ALCOL;
	} else if (carburante == VALORE_CHEROSENE) {
		prezzo_carburante = PREZZO_CHEROSENE;
	}

	if(prezzo_carburante != 0.0) {
		fmt.Print("litri? ");
		fmt.Scan(&litri);

		fmt.Println(prezzo_carburante*litri);
	} else {
		fmt.Println("valore non valido");
	}
}