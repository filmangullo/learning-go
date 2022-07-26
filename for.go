package main

import "fmt"

func main() {
    counter := 1;

    for counter <= 10 {
        fmt.Println("Perulangan ke ", counter);
        counter++
    }

//     For with statement
//     INIT STATEMENT             POST STATEMENT
//       ||                         ||
//      \||/                       \||/
    for value := 1; value <= 10; value++ {
        fmt.Println("Value ke ", value)
    }

// using For
    slice := []string{"Nol", "Satu", "Dua", "Tiga",};

    for i := 0; i< len(slice); i++ {
        fmt.Println(slice[i]);
    }

// Using For Range
    days := []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "sabtu",};

    for i, day := range days {
        fmt.Println("Index ", i, "= ", day);
    }

//      Var jika tidak di gunaka bisa menggunaka _
    for _, day := range days {
        fmt.Println("Hari ", day);
    }
}