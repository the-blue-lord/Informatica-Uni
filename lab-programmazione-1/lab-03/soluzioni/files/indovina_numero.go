package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const MAX = 13;
	
	SEGRETO := rand.Intn(MAX) + 1;
	var n, i int;

	for i = 1; n != SEGRETO && i <= MAX/2; i++ {
		fmt.Print("indovina il numero: ");	
		fmt.Scan(&n);
		if(n < 1 || MAX < n) {
			fmt.Println("fuori intervallo!");
			i--;
			continue;
		}
	}

	if(n != SEGRETO) {
		fmt.Println("hai perso, il numero era", n);
		return;
	}

	fmt.Println("hai indovinato! N. di tentativi:", i);
}