package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var lower, upper int = 0, 12;
	fmt.Println("10 numeri casuali tra", lower, "e", upper);
	for i := 1; i <= 10; i++ {
		fmt.Print(rand.Intn(upper), " ");
	}
	fmt.Println();
	
	/*
		Completare la seguente parte del programma in modo che generi e stampi 10 numeri casuali compresi in un intervallo definito dall'utente.
	*/

	fmt.Println("Ti fornirò 10 numeri causali in un intervallo a tua scelta");
	fmt.Print("Dammi gli estremi inferiore e superiore dell'intervallo: ");
	fmt.Scan(&lower, &upper);

	for i := 0; i < 10; i++ {
		fmt.Print(rand.Intn(upper-lower+1)+lower, " ");
	}
	fmt.Println();
}