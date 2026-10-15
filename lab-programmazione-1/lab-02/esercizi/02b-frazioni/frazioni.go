package main

import "fmt"

func main() {
	var n1, d1, n2, d2 int;

	fmt.Print("num e den fraz 1: ");
	fmt.Scan(&n1, &d1);
	fmt.Print("num e den fraz 2: ");
	fmt.Scan(&n2, &d2);

	frazioni_equivalenti := n1 * d2 == n2 * d1;

	if(!frazioni_equivalenti) {
		fmt.Print("non ");
	}

	fmt.Println("equivalenti");
}