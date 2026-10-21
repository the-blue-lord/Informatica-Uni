package main

import "fmt"

func main() {
	/* 	 
		Scrivere un programma che accetta come input un voto compreso tra 0 e 100. In caso di un valore del voto non valido stampa a schermo "voto non valido", altrimenti se l'input è valido trasforma il voto numerico in un voto letterale secondo la sguente tabella.

		+--------------+-----------+
		|     voto     |  lettera  |
		+--------------+-----------+
		|   [90, 100]  |     A     |
		+--------------+-----------+
		|   [80, 90)   |     B     |
		+--------------+-----------+
		|   [70, 80)   |     C     |
		+--------------+-----------+
		|   [60, 70)   |     D     |
		+--------------+-----------+
		|   [0, 60)    |     F     |
		+--------------+-----------+
	*/

	var voto int;
	fmt.Print("voto? ");
	fmt.Scan(&voto);

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido");
		return; //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A");
	} else if voto >= 80 {
		fmt.Println("B");
	} else if voto >= 70 {
		fmt.Println("C");
	} else if voto >= 60 {
		fmt.Println("D");
	} else {
		fmt.Println("F");
	}
}
/* DOMANDE */
/*
- D1. Per quali valori è vera la condizione del primo if? (quello subito dopo la Scan)
- R1. Per tutti i voti inferiori a 0 oppure maggiori a 100, ovvero per tutti i voti compresi tra 0 e 100.

- D2. Per quali valori viene eseguita la seguente istruzione?
    fmt.Println("B")
- R2. Per tutti i voti compresi tra 80 e 89

- D3. Giustificate la vostra risposta alla domanda D2
- R3. Esclusi i valori maggiori o uguali di 90 dall'if precedente, la condizione del secondo if viene controllata solo se il voto e minore o uguale a 89, viene aggiunta però proprio dalla condizione dell'if che il voto deve essere maggiore o uguale a 80, dunque per eseguire l'istruzione fmt.Println("B"), presente nel blocco del secondo if della catena, il voto deve essere compreso tra 80 e 89.

- D4. Il programma è impostato correttamente o sarebbe stato (più) corretto scrivere invece così?
 } else if voto >= 80 && voto < 90 {
- R4. Sarebbe stato assolutamente indifferente.

- D5. Perché?
- R5. Se viene controlalta la condizione del secondo if (voto >= 80 && voto < 90), vuol dire che la condizione del primo (oto >= 90) è riusltata falsa, dunque sarà sempre vero che (voto < 90). Di conseguenza in questo particolre caso, in questa condizione così costruita (voto >= 80 && voto < 90) è uguale a (voto >= 80).
*/