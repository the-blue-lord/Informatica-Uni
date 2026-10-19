package main

import "fmt"

func main() {
	/*
		Scrivere un programma che prende come input due interi uno per selezionare una figura (rettangolo e trinagolo) e uno per selezionare la formula richiesta (quella per il perimetro o quella per l'area). In base agli input forniti stampare a schermo la formula della proprietà indicata per la figura richiesta.
	*/
	const (
		RETTANGOLO = 1
		TRIANGOLO  = 2
		PERIMETRO  = 1
		AREA       = 2
	);
	var figura, operazione int;

	fmt.Println("di che figura si tratta?");
	fmt.Print("rettangolo (1) o triangolo (2)? ");
	fmt.Scan(&figura);
	fmt.Println("cosa vuoi calcolare?");
	fmt.Print("perimetro (1) o area (2)? ");
	fmt.Scan(&operazione);

	if figura == RETTANGOLO {
		if operazione == PERIMETRO {
			fmt.Println("formula = (base + altezza) * 2");
		} else { //operazione == AREA
			fmt.Println("formula = base * altezza");
		}
	} else { //figura == TRIANGOLO
		if operazione == PERIMETRO {
			fmt.Println("formula = lato1 + lato2 + lato3 ");
		} else { //operazione == AREA
			fmt.Println("formula = (base * altezza)/2");
		}
	}
}
/* DOMANDE */
/*
- D1. Quanti diversi messaggi da stampare ci sono?
- R1. Quattro messaggi, due formule per due figure: 2*2=4.

- D2. Sarebbe possibile ridurre il numero di if/else? Perché?
- R2. Sarebbe possibile ridurre il numero di if/else da 6 a 4 unendo el condizioni tramite condizione e if else, eliminando gli if annidati.
*/