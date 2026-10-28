package main

import "fmt"

func main() {
    /*
		Richiedi K=5 interi all'utente e stampane la somma.
    */
    const K = 5;
	var n int;
	somma := 0;
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ");
		fmt.Scan(&n);
		somma += n;  // somma = somma + n
	}
	fmt.Println(somma);
}