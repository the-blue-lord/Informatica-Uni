package main;

import ("fmt"; "math");

func main() {
	var x1, y1, x2, y2 float64;

	fmt.Print("x e y del primo punto: ");
	fmt.Scan(&x1, &y1);
	fmt.Print("x e y del secondo punto: ");
	fmt.Scan(&x2, &y2);

	dx := x1 - x2;
	dy := y1 - y2;

	distanza := math.Sqrt(dx*dx + dy*dy);

	fmt.Println("La distanza tra i due punti è:", distanza);
}
