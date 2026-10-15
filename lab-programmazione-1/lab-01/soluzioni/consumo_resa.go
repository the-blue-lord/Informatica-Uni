package main;

import "fmt";

func main() {
	var distanza, litri, consumo_medio, resa_media float64;

	fmt.Print("distanza percorsa (in km): ");
	fmt.Scan(&distanza);
	fmt.Print("quantità di carburante utilizzata (in l): ");
	fmt.Scan(&litri);

	consumo_medio, resa_media = 0, 0;

	if(distanza != 0 && litri != 0) {
		consumo_medio = litri/distanza;
		resa_media = distanza/litri;
	}

	fmt.Println("consumo medio:", consumo_medio, "l/km");
	fmt.Println("resa media:", resa_media, "km/l");
}
