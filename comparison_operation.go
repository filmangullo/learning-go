package main

import "fmt"

func main() {
    var name1 = "Filman";
    var name2 = "Filman";

    var value1 = 10;
    var value2 = 20;

    var comparisonName bool = name1 == name2;

    fmt.Println(comparisonName);
    fmt.Println(value1 > value2);
    fmt.Println(value1 < value2);
    fmt.Println(value1 == value2);
    fmt.Println(value1 != value2);
}