package main

import "fmt"

func main() {
	const K = 5;
	var n int;

	for i := 0; i < K; i++ {
		fmt.Print("intero: ");
		fmt.Scan(&n);
		fmt.Println("doppio:", 2*n);
		fmt.Println();
	}
}