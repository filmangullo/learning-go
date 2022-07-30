package main

import "fmt"

func getIdentity() (name string, nationality string, address string ){
    name = "Filimantaptius"
    nationality = "Indonesian"
    address = "Medan Perjuangan - Medan - Indonesia"

    return
}

func main() {
    name,nat,add := getIdentity()

    fmt.Println(name)
    fmt.Println(nat)
    fmt.Println(add)
}