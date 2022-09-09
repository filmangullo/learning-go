package main

import "fmt"

// Recover berfungsi menangkap sebuah message pada errored
// Recover berfungsi jika ada error maka program tetap berjalan.errored

// Fungsi recover sebenarnya untuk menangkap message dari panic sehigga Exception error tidak terjadi

func endApp(){
    message := recover();
    if message != nil {
        fmt.Println("Terjadi Error :",message);
    }

    fmt.Println("Aplikasi Selesai");

}

func runApp(error bool){
    defer endApp();
    if error {
        panic("Aplikasi error");
    }
    fmt.Println("Aplikasi Berjalan");
}

func main() {
    runApp(true);
    fmt.Println("Aplikasi Tetap Berjalan");
}