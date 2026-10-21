package main

import "fmt"

func main() {
	const (
		ALIQUOTA_INFERIORE = 0.1;
		ALIQUOTA_SUPERIORE = 0.24;
		SOGLIA_NON_CONIUGATO = 32000;
		SOGLIA_CONIUGATO = 64000;
		VALORE_PER_TRUE = "t" 
	);

	var r int;
	var cs string;
	var c bool;

	fmt.Print("reddito? ");
	fmt.Scan(&r);
	fmt.Print("coniugato? (true/false) ");
	fmt.Scan(&sc);

	if(cs == VALORE_PER_TRUE) {
		c = true;
	}

	fmt.Print("tasse per ")
	if(!c) {
		fmt.Print("non ");
	}
	fmt.Print("coniugato con reddito ", r, ": ");

	if(c && r <= SOGLIA_CONIUGATO || !c && <= SOGLIA_NON_CONIUGATO) {
		fmt.Println(r*ALIQUOTA_INFERIORE);
	}
	else {
		fmt.Println(r*ALIQUOTA_SUPERIORE);
	}
}