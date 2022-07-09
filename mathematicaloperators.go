package main

import "fmt"

func main() {
//  mathematic operation
    var a = 10;
    var b = 20;

    var c = a + b;
    fmt.Println(c);
    fmt.Println(a * b);

//  Augmented Asignment
    var i = 10;
    i += 10;
    fmt.Println(i);

//  Unary operation
    var married = true;
    var numberPlus = 10;
    var numberMinus = 10;
    var valuePostive = 100;
    var valueNegative = -100;

    numberPlus++;
    numberMinus--;

    fmt.Println(!married);
    fmt.Println(numberPlus);
    fmt.Println(numberMinus);
    fmt.Println(valuePostive);
    fmt.Println(valueNegative);
}