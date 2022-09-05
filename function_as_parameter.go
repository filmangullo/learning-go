package main

import "fmt"

// With declaration type data
type Filter func(string) string;
func sayHelloWithSpam(name string, filter Filter)  {
    nameFiltered := filter(name)
    fmt.Println("Hello", nameFiltered);
}

// Without declaration type data
func sayHelloWithFilter(name string, filter func(string) string){
    nameFiltered := filter(name)
    fmt.Println("Hello", nameFiltered);
}

func spamFilter(name string) string {
    if name == "Anjing" {
        return "******";
    }else {
        return name;
    }
}



func main() {
    sayHelloWithFilter("Filman", spamFilter);
    sayHelloWithFilter("Anjing", spamFilter);

    sayHelloWithSpam("Anjing", spamFilter);

}