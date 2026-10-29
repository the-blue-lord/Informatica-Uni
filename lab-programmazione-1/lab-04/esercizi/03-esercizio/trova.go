/*

Scrivere un programma trova.go che chiede ("un carattere: ") e legge da standard input un carattere (runa) e chiede ("una parola: ") e legge, sempre da standard input una stringa e stampa la posizione (0-based) del carattere nella stringa (la prima volta che appare), o -1 se il carattere non c'è.

*/


package main

import "fmt"

func main() {
	var char rune;
	var tmp, str string;

	fmt.Print("un carattere: ");
	fmt.Scan(&tmp);
	fmt.Print("una parola: ");
	fmt.Scan(&str);

	for _, r := range(tmp) {
		char = r;
		break;
	}

	for i, r := range(str) {
		if(r == char) {
			fmt.Println(i);
			return;
		}
	}

	fmt.Println(-1);
}