package main

import "fmt"

func main() {
//  Conversion type data integer
    var nilai32 int32 = 31768;
    var nilai64 int64 = int64(nilai32);

    var nilai16 int16 = int16(nilai64);

    fmt.Println(nilai16);

//  Conversion type integer to String
    var name = "Filimantaptius";
    var i = name[1];
    var iString string = string(i);

    fmt.Println(name);
    fmt.Println(iString);
}