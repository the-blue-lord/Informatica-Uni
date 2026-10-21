package main

import "fmt"

func main() {
	const COMPONENTE_1 = "caricamento acqua";
	const COMPONENTE_2 = "scarico acqua";
	const COMPONENTE_3 = "sistema di riscaldamento";

	var c_1, c_2, c_3 int;

	fmt.Print("componente 1, guasto rilevato? (0 per ok): ");
	fmt.Scan(&c_1);
	fmt.Print("componente 2, guasto rilevato? (0 per ok): ");
	fmt.Scan(&c_2);
	fmt.Print("componente 3, guasto rilevato? (0 per ok): ");
	fmt.Scan(&c_3);

	if(c_1 != 0 || c_2 != 0 || c_3 != 0) {
		fmt.Println("guasti rilevati:");
	} else {
		fmt.Println("nessun guasto rilevato");
	}

	if(c_1 != 0) {
		fmt.Println(COMPONENTE_1);
	}
	if(c_2 != 0) {
		fmt.Println(COMPONENTE_2);
	}
	if(c_3 != 0) {
		fmt.Println(COMPONENTE_3);
	}
}