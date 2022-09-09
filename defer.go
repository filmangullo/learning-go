package main

import "fmt"

// Defer berfungsi jika ada function tetap di pakai jika func sebelumnya telah selesai di jalankan.
func logging(){
    fmt.Println("Selesai Memanggil Functions")
}

func runApplication(value int){
    defer logging();
    // jika terjadi error setelah defer maka function logging tetap di pakai
    // Itu fungsi defer
    fmt.Println("Run runApplication");
}

func main() {
    runApplication(10);
}