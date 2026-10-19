package main

import "fmt"

func main() {
	var ora int;

	fmt.Print("ora: ");
	fmt.Scan(&ora);

	if(ora < 0 || 23 < ora) {
		fmt.Println("orario non valido");
		return;
	}

	if(ora <= 6) {
		fmt.Println("notte");
	} else if(ora <= 13) {
		fmt.Println("mattino");
	} else if(ora <= 18) {
		fmt.Println("pomeriggio");
	} else {
		fmt.Println("sera");
	}
}