package main

import "fmt"

func main() {
	var n, a, sum int;

	fmt.Print("un numero intero positivo: ");
	fmt.Scan(&n);
	a = n;

	for ; n != 0; n /= 10 {
		sum += n%10;
	}

	fmt.Println("la somma delle cifre di", a, "è", sum);
}