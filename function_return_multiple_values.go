package main

import "fmt"

func getFullname() (string, string, int) {
    return "Filman", "Gulo", 170
}

func main(){
//  you can use _ for ignore variable
    firstName, lastName, _ := getFullname();
    fmt.Println(firstName);
    fmt.Println(lastName);


}