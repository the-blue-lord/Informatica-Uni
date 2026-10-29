/*

Scrivere un programma lunghezza_tot.go che chiede ("lunghezza totale: ") e legge da standard input un int totLen, poi legge (senza stampare richieste) una sequenza di stringhe di caratteri ASCII, ne somma le lunghezze fino a raggiungere (o superare) totLen. 
Raggiunto totLen, il programma stampa, separate da uno spazio, la somma delle lunghezze e la concatenazione delle stringhe lette e termina

*/

package main

import "fmt"

func main() {
	fmt.Print("lunghezza totale: ");

	var totLen int;
	fmt.Scan(&totLen);

	lun := 0;
	res := "";

	for str := ""; lun < totLen; res, lun = res + str, lun + len(str) {
		fmt.Scan(&str);
	}

	fmt.Println(lun, res);
}