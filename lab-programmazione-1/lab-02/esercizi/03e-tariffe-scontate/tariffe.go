package main

import "fmt"

func main() {
	const VALORE_PER_TRUE = "t";



	var e int;
	var s string;

	fmt.Print("età: ");
	fmt.Scan(&e);
	fmt.Print("studente? (t/f): ");
	fmt.Scan(&s);

	if(s == VALORE_PER_TRUE) {
		if(e >= 18) {
			e = 10;
		}
	}

	var tariffa float64;

	if(9 <= e && e < 18) {
		tariffa = 5;
	} else if(18 <= e && e < 25) {
		tariffa = 7.5;
	} else if(25 <= e && e < 65) {
		tariffa = 10;
	} else if(65 <= e) {
		tariffa = 7.5;
	}

	fmt.Println("ingresso", tariffa, "euro");
}