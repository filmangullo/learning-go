package main

import (
    "errors"
    "fmt"
)

func Pembagian (nilai int, pembagi int) (int, error) {
    if pembagi == 0 {
        return 0, errors.New("Pembagi must be a positive")
    } else {
        result := nilai / pembagi;
        return result, nil;
    }
}

func main() {
    hasil, err := Pembagian (100, 0);
    if err == nil {
        fmt.Println("Hasil : ", hasil);
    }else {
        fmt.Println(hasil, err.Error());
    }
}