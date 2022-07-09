package main

import "fmt"

func main() {
    type Data string;
    type Value int8;

    var fullName Data = "Filimantaptius Gulo";
    var english Value = 80;
    var mathematic Value = 90;

    fmt.Println(fullName);
    fmt.Println(english);
    fmt.Println(mathematic);
}