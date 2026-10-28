package main

import ("fmt";"math/rand")

func main() {
	var n, k int;

	for i := 0; i < 10; i++ {
		n = rand.Intn(11);
		if(n%2 == 0) {
				k++;
		}
		fmt.Print(n, " ");
	}
	
	fmt.Println("\nNumeri pari generati:", k);
}