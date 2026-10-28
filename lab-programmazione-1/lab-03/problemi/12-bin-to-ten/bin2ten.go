package main

import "fmt"

func main() {
	var n, res int;

	fmt.Print("un numero binario: ");
	fmt.Scan(&n);

	is_negative := n < 0;

	if(is_negative) {
		n = -n;
	}

	for pow := 1; n != 0; pow, n = 2*pow, n/10 {
		c := n%10;

		if(c == 1) {
			res += pow;
		} else if(n%10 >= 2) {
			fmt.Println("dato non valido");
			return;
		}
	}

	if(is_negative) {
		fmt.Print("-");
	}
	fmt.Println(res);
}