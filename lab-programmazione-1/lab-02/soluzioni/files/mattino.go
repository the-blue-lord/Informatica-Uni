package main

import "fmt"

func main() {
	var i int;

	fmt.Print("scrivi un orario (int): ");
	fmt.Scan(&i);

	if(i >= 6 && i < 13) {
		fmt.Println("mattino");
	}
}