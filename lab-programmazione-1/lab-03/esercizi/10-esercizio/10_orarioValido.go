package main

import "fmt"

func main() {
	var ore, min int = -1, -1;

	for (ore < 0 || 23 < ore || min < 0 || 59 < min) {
		fmt.Print("ore e minuti: ");
		fmt.Scan(&ore, &min);
	}

	fmt.Println("orario valido");
}
