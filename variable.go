package main

import "fmt"

func main() {
//  Var with initial type
    var name string
    var height int16

// Variable without var declaration
    city := "Medan";

//  Multiple variable declarations
    var (
        firstName = "Filimantaptius";
        lastName = "Gulo";
    )
    fmt.Println(firstName);
    fmt.Println(lastName);

//  Var String and Manipulation
    name = "Filimantaptius Gulo";
    fmt.Println(name);

    name = "Filman";
    fmt.Println(name);

    city = "Medan Perjuangan"
    fmt.Println(city);


//  Var Int
    height = 17;
    fmt.Println(height);

//  Var with initial type
    var nationality = "Indonesian";
    fmt.Println(nationality);

    var age = 25
    fmt.Println(age)
}