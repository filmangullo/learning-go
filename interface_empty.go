package main

import "fmt"

// Golang bukan bahasa pemrograman yang berorientasi objek
// Interface Kosong adalah interface yang tidak memiliki deklarasi method satupun, hai ini membuat secara
// otomatis semua tipe data akan menjadi implementasinya.

func Ups(i int) interface{} {
    if i == 1 {
        return 1;
    } else if i == 2 {
        return true;
    } else {
        return "Ups";
    }
}

func main() {
    var data interface{} = Ups(1);
    fmt.Println(data);
}