package main

import "fmt"

// Panic berfungsi untuk memberhentikan aplikasi untuk sebuah error.
func endApp(){
    fmt.Println("Aplikasi Selesai")
}

func runApp(error bool){
    defer endApp();
    if error{
    // Memberi perintah error
        panic("Aplikasi Error");
    }
}

func main() {
    runApp(true);
}