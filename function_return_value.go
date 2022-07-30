package main

import "fmt"

func getHeightInc(cm int) float64 {
    var centimeter float64 = float64(cm)
    result := centimeter * 0.393701;
    return result
}

func getHello(name string) string {
    return "Hello " + name;
}

func main() {
    greeting := getHello("Filman");
    fmt.Println(greeting);
    height := getHeightInc(170);
    fmt.Println(height);
}