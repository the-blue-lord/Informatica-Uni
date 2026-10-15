package main

import "fmt"

func main() {
    var x, z int;

    x = 3;

    z = (z = 5);

    fmt.Println(x);
}
