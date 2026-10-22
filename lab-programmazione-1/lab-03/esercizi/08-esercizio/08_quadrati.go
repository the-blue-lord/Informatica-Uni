package main

import "fmt"

func main() {
	var n int;

	fmt.Scan(&n);

	for i := 1; i*i <= n; i++ {
		fmt.Println(i*i);
	}
}