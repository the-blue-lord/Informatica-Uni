package main

import "fmt"

func main() {
	var c rune;
	var sum int;
	for (c != '.') {
		fmt.Scanf("%c", &c);
		if(c == '.' || c == '\n') {
			continue;
		}

		fmt.Print(string(c));
		if('a' <= c && c <= 'z') {
			fmt.Print(" è la ", c - 'a' + 1, "^a\n");
			continue;
		}

		fmt.Print(" - ");

		if('0' <= c && c <= '9') {
			i := int(c) - '0'
			sum += i;
			fmt.Println(i);
			continue;
		}

		fmt.Println("altro");
	}
	fmt.Println("somma:", sum);
	fmt.Println("bye");
}