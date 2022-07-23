package main

import "fmt"

func main() {
    name := "Gulo";

    if name == "Filman" {
        fmt.Println("Hello ", name);
    } else if name == "Gulo" {
        fmt.Println("Hello", name);
    } else {
        fmt.Println("Hi, who are you ?");
    }

// Normal Statement
    var length = len(name);
    if length > 5 {
        fmt.Println("Nama Telalu Panjang");
    } else {
        fmt.Println("Panjang Nama Cocok");
    }

//  if with Short Statement
    if length := len(name) ;length > 5 {
        fmt.Println("Nama Telalu Panjang");
    } else {
        fmt.Println("Panjang Nama Cocok");
    }

}