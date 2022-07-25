package main

import "fmt"

func main() {
    name := "Filman";

//  Normal Switch
    switch name {
        case "Filman":
            fmt.Println("Ini adalah nama panggilan");
        case "Filimantaptius Gulo":
            fmt.Println("Ini adalah Nama Panjang");
        default:
            fmt.Println("Apakah ini sebuah Nama");
    }

//  switch with Short Statement
    switch length := len(name); length > 6 {
        case true:
            fmt.Println("Nama Telalu Panjang");
        case false:
            fmt.Println("Nama Sudah Benar");
    }

//  Switch without expression
    length := len(name);
    switch {
        case length > 10:
        fmt.Println("Nama Telalu Panjang");
        case length < 5:
        fmt.Println("Nama Terlalu Pendek");
        default:
        fmt.Println("Nama Sudah Normal");
    }
}