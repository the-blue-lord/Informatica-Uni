package main

import "fmt"

func main() {
	var str string;

	fmt.Print("Inserisci un carattere: ");
	fmt.Scan(&str);

	var c rune;
	for _, r := range(str) {
		c = r;
		break;
	}
	fmt.Println(string(c));

	fmt.Println(string(c-1), string(c), string(c+1));

	if(c >= 'A' && c <= 'L') {
		fmt.Println("A-L");
	} else {
		fmt.Println("altro");
	}

	fmt.Print("\nInserisci una stringa: ");
	fmt.Scan(&str);

	for _, r := range(str) {
		fmt.Println(string(r));
	}
}
