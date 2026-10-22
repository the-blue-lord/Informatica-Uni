package main

import "fmt"

func main() {
	/*
		Prendi un input intero e usalo come valore di riferimento, successivamente continua a chiedere un nuovo input intero, aggiornare il valore di riferimento a un nuovo valore rappresentato con la concatenazione delle ultime due cifre del vecchio valore di riferimento e le ultime due cifre dell'input ricevuto e stampare questo nuovo valore di riferimento fino a quando l'utente non inserisce in successione due valori le cui ultime due cifre sono rappresentate rispettivamente da 1 e 0.
	*/
	var cifre_totali, cifre_iniziali, cifre_finali int;
	fmt.Scan(&cifre_finali);
	for cifre_totali != 100 {
		cifre_iniziali = cifre_finali;
		fmt.Scan(&cifre_finali);
		cifre_totali = (cifre_iniziali%100)*100 + cifre_finali%100;
		fmt.Println(cifre_totali);
	}
}