package main

import ("fmt"; "math/rand")

func main() {
	const TARGET = 50;
	const MAX = 10;

	var sum, cnt int;

	for sum < TARGET {
		cnt++;
		rnd := rand.Intn(MAX)+1;
		sum += rnd;
		fmt.Print(rnd, " ");
	}

	fmt.Println();
	fmt.Println(cnt, "numeri casuali");
	fmt.Println("somma:", sum);
	fmt.Println("media:", float64(sum)/float64(cnt));
}