package main

import "fmt"

func random() interface{} {
    return 100;
}

func main() {
    var result interface{} = random();
//     var resultString string = result.(string);

//     fmt.Println(resultString);

//  Lebih Aman
//  Diharapkan untuk Type Assertion menggunakan switch untuk menghindari Panic
    switch value := result.(type) {
    case string:
        fmt.Println("value ", value, " is string");
    case int:
        fmt.Println("value ", value, " is integer");
    default:
        fmt.Println("unknown");
    }

}
