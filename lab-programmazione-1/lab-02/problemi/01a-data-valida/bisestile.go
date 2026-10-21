package main

import "fmt"

func main() {
	var i int;

	fmt.Print("aaaa? ");
	fmt.Scan(&i);

	if(i%400 != 0 && (i%4 != 0 || i%100 == 0)) {
		fmt.Print("non ");
	}

	fmt.Println("bisestile");
}