package main

import "fmt"

type Address struct {
    City, Province, Country string
}

type Position struct {
    City, Province, Country string
}


func main() {
    // Past By Value
    address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
    address2 := address1;

    address2.City = "Nias";

    fmt.Println(address1);
    fmt.Println(address2);

    // Past By Reference
    position1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
    position2 := &position1;

    position2.City = "Binjai";

    fmt.Println(position1);
    fmt.Println(position2);

    // Opertor bintang
    var location1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
    var location2 *Address = &location1;
    var location3 *Address = &location1;

    location2.City = "Binjai";

    *location2 = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"};

    fmt.Println(location1);
    fmt.Println(location2);
    fmt.Println(location3);
    // siapa pun yang mengacu pada memory yang sama akan berubah lihat pada kasus location

    // Function New
    var place1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
    var place2 *Address = &place1;
    var place3 *Address = &place1;

    location2.City = "Binjai";

    *place2 = Address{"Jakarta Selatan", "DKI Jakarta", "Indonesia"};

    fmt.Println(place1);
    fmt.Println(place2);
    fmt.Println(place3);

    var place4 *Address = new(Address);
    fmt.Println(place4); // Ketika baru dibuat dan masih kosong
    place4.City = "Malang";
    fmt.Println(place4); // Ketika sudah di isi dan sudah berisi
}

// OPERATOR & PADA POINTER
// Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunkan membuat
// operator & diikuti dengan nama variable nya

// OPERATOR * PADA POINTER
// Jika kita ingin mengubah selurh data variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
// Semua variable yang mengacu ke data yang sama tidak akan berubah.
// Ketika kita mengubah pointer, maka yang berubah hanya variablenya saja.

// FUNCTION NEW
// Function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal.
