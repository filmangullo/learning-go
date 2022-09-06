package main

import "fmt"

// Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
// Dengan menerapkan konsep tersebut, kita bisa membuat fungsi di dalam fungsi, atau bahkan membuat fungsi yang mengembalikan fungsi.

// Closure merupakan anonymous function atau fungsi tanpa nama.
// Biasa dimanfaatkan untuk membungkus suatu proses yang hanya dipakai sekali atau dipakai pada blok tertentu saja.

func main() {
    name := "Filman";
    counter := 0;


    increment := func() {
        // Handle supaya name tidak berubah.
        name := "Gulo";
        fmt.Printf(name);
        // Counter dapat di akses dari dalam function Anonymous
        counter++;
    }

    increment();

    fmt.Println(counter);
    fmt.Println(name);
}