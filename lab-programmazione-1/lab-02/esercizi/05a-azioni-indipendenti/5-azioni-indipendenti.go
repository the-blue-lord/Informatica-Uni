package main

import "fmt"

func main() {
	/* 
		Scrivi un programma che legge un intero. Se é divisibile per 3 stampa "Frizz". Indipendentemente dalla condizione precedente se é divisibile per 5 stampa inoltre "Buzz".
	*/

	var num int;
	fmt.Print("numero? ");
	fmt.Scan(&num);

	if num%3 == 0 {
		fmt.Println("Fizz");
	}
	if num%5 == 0 {
		fmt.Println("Buzz");
	}
}
/* DOMANDE */
/*
- D1. Se al posto del secondo if ci fosse un else if (legato al primo if), il programma si comporterebbe nello stesso modo? Dareste le stesse specifiche?
- R1. No. Direi che deve stampare "Buzz" se é divisibile per 5 ma non per 3 e non scriverei che la seconda condizine è indipendente dalla seconda.

- D2. Perché? Spiegate la vostra risposta alla domanda D1
- R2. Perché il secondo blocco condizionale verrebbe eseguito solo se la condizione del primo è falsa, mentre nel programma originale il blocco condizione viene sempre valutato.

- D3. Cosa stampa il programma per ciascuno dei seguenti input: 8, 9 10, 15?
- R3
   8: ""
   9: "Frizz"
  10: "Buzz"
  15: "Frizz\nBuzz"
 

*/