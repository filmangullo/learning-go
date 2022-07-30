package main

import "fmt"

func main() {
//  With declaration Value
    person := map[string]string{
        "name": "Filman",
        "nationality": "Indonesian",
    }

    person["title"] = "Mr.";

    fmt.Println(person);
    fmt.Println(person["name"]);
    fmt.Println(person["nationality"]);

//  Count data in Map
    fmt.Println(len(person));

//  Make new Map
    var book map[string]string = make(map[string]string);

    book["title"] = "Belajar Golang";
    book["author"] = "Filman";
    book["ups"] = "Salah";

    fmt.Println(book);

//  Delete data in map
    delete(book, "ups");

    fmt.Println(book);


}