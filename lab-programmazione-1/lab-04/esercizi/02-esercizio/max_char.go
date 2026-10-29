/*

Scrivere un programma max_char.go che legge (senza stampare richieste) da standard input una sequenza di 5 caratteri ASCII (byte), non separati da spazi, da a capo o da altro, e stampa il maggiore in ordine lessicografico (cioè con il codice ASCII più alto).

*/

package main

import "fmt"

func main() {
	var str string;

	var max byte;

	fmt.Scan(&str);

	for _, r := range(str) {
		if(int(r) > int(max)) {
			max = byte(r);
		}
	}

	fmt.Println(string(max));
}