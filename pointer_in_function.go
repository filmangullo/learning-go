package main

import "fmt"

type Address struct {
    City, Province, Country string
}

// Untuk menjadikan sebuah parameter sebagai pointer, kota bisa menggunakan opertor * di parameternya.
func changeAddressToIndonesia(address *Address) {
    address.Country = "Indonesia";
}

func main() {
    address := Address{"Medan", "Sumatera Utara", ""}

    changeAddressToIndonesia(&address);

    fmt.Println(address)

// other
    address1 := Address{"Binjai", "Sumatera Utara", ""}
    var addressPointer *Address = &address1;

    changeAddressToIndonesia(addressPointer);

    fmt.Println(addressPointer)
}