package main

import "fmt"

type Customer struct {
    Name string;
    Address string;
    Age int;
    Married bool;
}
func (customer  Customer) sayHelloFrom(name string) {
    fmt.Println("Hello, ", customer.Name, " Welcome. By : ", name )
}

func (user  Customer) sayHelloTo() {
    fmt.Println("Hello, Welcome back : ", user.Name )
}

func main() {
    var user Customer;

    user.Name = "Filman";
    user.Address = "Medan - Indonesian";
    user.Age = 25;

    user.sayHelloFrom("Super Admin");
    user.sayHelloTo();
}