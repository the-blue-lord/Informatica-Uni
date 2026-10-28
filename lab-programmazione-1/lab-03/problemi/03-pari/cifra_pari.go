package main

import "fmt"

func main() {
	var a, i int;

	fmt.Print("un numero intero positivo: ");
	fmt.Scan(&a);

	for i = 1; a != 0; i++ {
		if(a%2 == 0) {
			break;
		}
		a /= 10;
	}

	fmt.Print("la prima cifra pari (", a%10, ") è in posizione ", i, "\n");
}