package main;

import "fmt";

func main() {
	const SOMMA_ANGOLI_TRIANGOLO = 180;

	var a1, a2 float64;

	fmt.Print("angolo uno e due: ");
	fmt.Scan(&a1, &a2);

	a3 := SOMMA_ANGOLI_TRIANGOLO - a1 - a2;

	fmt.Println("ampiezza terzo angolo: ", a3);
}
