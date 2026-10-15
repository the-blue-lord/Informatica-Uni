package main;

import "fmt";

func main() {
	const GIORNI_PER_ANNO = 365;
	const GIORNI_PER_SETTIMANA = 7;

	var input, anni, settimane, giorni int;

	fmt.Print("numero di giorni da convertire? ");
	fmt.Scan(&input);

	giorni = input;

	anni = giorni / GIORNI_PER_ANNO;
	giorni = giorni % GIORNI_PER_ANNO;

	settimane = giorni / GIORNI_PER_SETTIMANA;
	giorni = giorni % GIORNI_PER_SETTIMANA;

	fmt.Println(input, "giorni =", anni, "anni +", settimane, "settimane +", giorni, "giorni");
}
