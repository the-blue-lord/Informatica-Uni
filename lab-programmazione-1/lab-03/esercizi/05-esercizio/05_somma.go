package main

import ("fmt"; "math/rand")

func main() {
	const K = 5;
	var sum int;

	for i:=0; i < K; i++ {
		r := rand.Intn(10) + 1;
		fmt.Println(r);
		sum += r;
	}

	fmt.Println("la somma è", sum);
}