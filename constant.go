package main

import "fmt"

func main() {
//  Singel Constants
    const FIRSTNAME string = "Filimantaptius";
    const LASTNAME = "Gulo"
//  Multiple Constants
    const (
        ADDRESS_1 string = "Jln. Menuju Programming";
        ADDRESS_2 = "Jln. Menuju Keberhasilan";
        NUMBER = 1000;
    );

    fmt.Println(FIRSTNAME);
    fmt.Println(LASTNAME);

    fmt.Println(ADDRESS_1);
    fmt.Println(ADDRESS_2);
    fmt.Println(NUMBER);

}