package main;

import "fmt";

func main() {
	var v1, v2 int;

	fmt.Print("val1 e val2 (int): ");
	_, err := fmt.Scan(&v1, &v2);

	fmt.Println("errore:", err);

	fmt.Println("prima dello scambio:", v1, v2);
	v1, v2 = v2, v1;

	fmt.Println("dopo lo scambio:", v1, v2);
}
