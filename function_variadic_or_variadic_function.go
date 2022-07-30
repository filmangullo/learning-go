package main

import "fmt"

func sumAll(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    total := sumAll(10, 20, 30, 15, 25, 11,)

    fmt.Println(total)

//  Make slice in functiion just add ... or titik 3 x
    slice := []int{10, 20, 30, 15, 26, 4}
    total = sumAll(slice...)
    fmt.Println(total)
}