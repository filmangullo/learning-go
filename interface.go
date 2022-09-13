package main

import "fmt"
// About
// Interface adalah tipe data Abstract, dia tidak memiliki implementation langsung.
// Sebuah interface berisikan definisi-definisi methods
// Biasanya interface digunkan sebagai kontrak


// Implementation
// Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut.
// Sehingga kita tidak perli mengimplementasikan interface secara manual
// Hal ini agak berbeda dengan bahasa pemograman lain yang yang ketika membuat interface, kita harus
// menyebutkan secara exsplisit akan menggunkan interfac mana

// Interface di gunakan untuk membuat kontrak

type HasName interface {
    GetName() string;
}

func SayWhat(hasname HasName){
    fmt.Println("What ", hasname.GetName());
}

type Person struct {
    Name string;
}

func (p Person) GetName() string {
    return p.Name;
}

type Animal struct {
    Name string;
}

func (a Animal) GetName() string {
    return a.Name;
}

func main() {
     var p Person;
     p.Name = "Filman";
     SayWhat(p);

     a := Animal {
        Name: "Tom",
     }
     SayWhat(a);
}