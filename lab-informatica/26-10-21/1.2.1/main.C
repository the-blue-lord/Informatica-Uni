#include <iostream>

int main(void) {
    int a, b;
    float c, d;
    std::cout << "Input\na: "; std::cin >> a;
    std::cout << "b: "; std::cin >> b;
    std::cout << "c: "; std::cin >> c;
    std::cout << "d: "; std::cin >> d;
    float x = a + c;
    float y = b * d;
    std::cout << "\nRisultati\n"
        << "x: \t" << x << "\n" 
        << "y: \t" << y << "\n"
        << "a/b: \t" << a / b << "\n"
        << "c/d: \t" << c / d << "\n"
        << "\nCon difre decimali\na/b: " << (float) a/b << "\n";
}