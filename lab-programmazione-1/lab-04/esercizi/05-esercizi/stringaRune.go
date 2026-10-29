/*

Scrivere un programma stringaRune.go che legge da standard input una stringa e la analizza runa per runa: stampa "nessuna cifra" se la stringa non contiene nessuna cifra, stampa invece "ha cifre" altrimenti.

*/

package main

import "fmt"

func main() {
	var str string;

	fmt.Scan(&str);

	for  _, r := range(str) {
		if('0' <= r && r <= '9') {
			fmt.Println("ha cifre");
			return;
		}
	}

	fmt.Println("nessuna cifra");
}