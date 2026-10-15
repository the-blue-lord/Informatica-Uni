package main

import "fmt"

func main() {
	const MINUTI_PER_ORA = 60;

	var ore, min int;

	fmt.Scan(&ore, &min);

	min += ore * MINUTI_PER_ORA;

	if(min >= 330 && min < 750) {
		fmt.Println("mattino");
	}
}