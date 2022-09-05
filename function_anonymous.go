package main

import "fmt"

type Blacklist func(string) bool;

func registerUser(name string, blacklist Blacklist){
    if blacklist(name) {
        fmt.Println("You are blocked");
    }else {
        fmt.Println("Welcome", name);
    }
}

func main() {
    // Anonymous function yaitu dimanan function bisa di buat di dalam fungsi pemanggilannya
    blacklist := func(name string) bool {
        return name == "admin";
    }
    // End Anonymous function

    registerUser("admin", blacklist);
    registerUser("Filman", blacklist);
    registerUser("Root", blacklist);
}