/*

Scrivere un programma stringaByte.go che legge da standard input una stringa e la analizza byte per byte: stampa "+" quando il carattere considerato è maggiore (ha un codice ASCII maggiore) del precedente, stampa "-" quando è minore e stampa "=" quando è uguale.

*/

package main

import "fmt"

func main() {
	var str string;
	fmt.Scan(&str);

	for i := 1; i < len(str); i++ {
		if(str[i] == str[i-1]) {
			fmt.Print("=");
			continue;
		}

		if(str[i] < str[i-1]) {
			fmt.Print("-");
			continue;
		}

		fmt.Print("+");
	}

	fmt.Println();
}