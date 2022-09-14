package main

import "fmt"

// Biasanya di dalam bahasa pemrograman lain, object yang belum di inisialisasi maka secara otomatis nilainya adalah "null" atau "nil"
// Nil sendiri hanya bisa digunkan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel

func NewMap(name string) map[string]string {
    if name == "" {
        return nil;
    } else {
        return map[string]string {
            "name" : name,
        }
    }
}

func main() {
    var post = NewMap("Gulo");
    fmt.Println(post);

    var person map[string] string = NewMap("Filman");
    if person == nil {
        fmt.Println("Data Kosong")
    } else {
        fmt.Println(person)
    }
}