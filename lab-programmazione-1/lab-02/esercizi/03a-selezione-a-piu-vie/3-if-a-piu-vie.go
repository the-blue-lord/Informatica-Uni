package main

import "fmt"

func main() {
	/*
		Scrivere un programma che che legga un intero e stampi a schermo se l'input è positivo, negativo o nullo.
	*/

	var num int;

	fmt.Print("un int: ");
	fmt.Scan(&num);

	if num > 0 {
		fmt.Println("positivo");
	} else if num == 0 {
		fmt.Println("nullo");
	} else { // < 0
		fmt.Println("negativo");
	}
}