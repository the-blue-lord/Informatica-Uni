package main;

import "fmt";

func main() {
	var m1, q1, m2, q2, x, y float64;

	fmt.Print("retta 1: m e q? ");
	fmt.Scan(&m1, &q1);
	fmt.Print("retta 2: m e q? ");
	fmt.Scan(&m2, &q2);

	x, y = 0, 0;

	if(m1 != m2) {
		x = (q2 - q1)/(m1 - m2);
		y = m1 * x + q1;
	}

	fmt.Print("intersezione in (", x, ",", y, ")\n");
}
