package main

import "fmt"

// Walaupun method akan menempel di struct, tetapi sebenarnya data struct yang di akses di dalam method adala pass by value.
// Sangat di rekomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method.

type Man struct {
    Name string
}

// Gunakan * (bintang) untuk mendakan pointer
func (man *Man) Married() {
    man.Name = "Mr. " + man.Name
}

func main() {
    name := Man{"Filman"}

    name.Married();

    fmt.Println(name.Name)
}