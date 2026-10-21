package main

import "fmt"

func main()  {
	var x, y int;
	fmt.Print("due int (x e y): ");
	fmt.Scan(&x, &y);

	fmt.Println(x != 0 && y/x  > 10);
}
