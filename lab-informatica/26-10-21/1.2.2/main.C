#include <iostream>

int main(void) {
    int i;

    std::cout << "Inserisci intero: ";
    std::cin >> i;

    int u = i%10;
    int d = i/10;

    int e = 16*d + u;

    std::cout << "Valore trasformato: " << e << "\n";
}