package main

import "fmt"

func main() {
//  Indirect declaration
    var names [3]string;
    names[0] = "Filman";
    names[1] = "Filimantaptius";
    names[2] = "Gulo";

    fmt.Println(names[0]);
    fmt.Println(names[1]);
    fmt.Println(names[2]);

//  Direct declaration
    var values = [3]int{
        90,
        95,
        80,
    };
    fmt.Println(values);

//  Array function len count lengt data in array
    var empty [10]string;

    fmt.Println(len(names));
    fmt.Println(len(empty));
}