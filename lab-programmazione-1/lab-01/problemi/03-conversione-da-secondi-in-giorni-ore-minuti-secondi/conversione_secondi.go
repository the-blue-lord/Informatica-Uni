package main;

import "fmt";

func main() {
	const SECONDI_PER_MINUTO = 60;
	const MINUTI_PER_ORA = 60;
	const ORE_PER_GIORNO = 24;

	var g, h, m, s int;

	fmt.Print("numero di secondi: ");
	fmt.Scan(&s);

	m = s / SECONDI_PER_MINUTO;
	s %= SECONDI_PER_MINUTO;

	h = m / MINUTI_PER_ORA;
	m %= MINUTI_PER_ORA;

	g = h / ORE_PER_GIORNO;
	h %= ORE_PER_GIORNO;

	fmt.Print("g:h:m:s = ", g, ":", h, ":", m, ":", s, "\n");
}
