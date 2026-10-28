package main

import "fmt"

func main() {
	var a, b int;

	fmt.Print("Inserisci due interi: ");
	fmt.Scan(&a, &b);

	if a <= 0 || b <= 0 {
		fmt.Println("Gli interi devono essere positivi non nulli");
		return;
	}

	for b != 0 {
		a, b = b, a%b;
	}

	fmt.Println("MCD calcolato con l'algoritmo di Euclide:", a);
}
