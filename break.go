package main

import "fmt"

func main() {

//  Break
    for i := 0; i < 10; i++ {
        if i == 5 {
            break;
        }
        fmt.Println("Perulangan Ke ", i);
    }

//  Continue
    for i := 0; i < 10; i++ {
        if i % 2 == 0 {
            continue
        }
        fmt.Println("Perulangan Ke ", i , "Ganjil");
    }
}