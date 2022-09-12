package main;

import "fmt";

// Struct adalah sebuat tempate data yang digunkan untuk menggabungkan nol atau lebih tipe data lainnya
// dalam satu kesatuan

// Standard golang use variable name with Uppercase to Lowercase

type Customer struct {
    Name string;
    Address string;
    Age int;
    Married bool;
}

func main() {
    // Struct Generals
    var user Customer;

    user.Name = "Filman";
    user.Address = "Medan - Indonesian";
    user.Age = 25;

    fmt.Println(user);
    fmt.Println(user.Name)

    // Struct Literals
    admin := Customer{
        Name: "Gulo",
        Address: "Medan - Indonesia",
        Age: 25,
    }
    fmt.Println(admin);
    fmt.Println(admin.Name)
}