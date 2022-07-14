package main

import "fmt"

func main() {
    var months = [...]string{
        "Jan", "Feb", "Mar", "Apr", "May","Jun", "Jul", "Aug", "Sep", "Oct", "Nov","Dec",
    };
    var sliceM1 = months[4:7]
    fmt.Println(sliceM1);
    fmt.Println(len(sliceM1));
    fmt.Println(cap(sliceM1));

//  Jika array tidak cukup untuk menyimpan data maka array baru dibuat
    var sliceM2 = append(sliceM1, "Hello");
    fmt.Println(sliceM2);
    sliceM2[1] = "June";
    fmt.Println(sliceM2);

    fmt.Println(sliceM1);
    fmt.Println(months);

    weekday := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri",};
    var sliceW1 = weekday[3:];
    fmt.Println(sliceW1);
//  Jika array masih cukup untuk menyimpan data maka array baru tidak dibuat
    var sliceW2 = append(sliceW1, "Sat");
    fmt.Println(sliceW2);
    sliceW2[1] = "Wednesday";
    fmt.Println(sliceW2);

    fmt.Println(sliceW1);
    fmt.Println(weekday);

//  Simple Program make slice
    newSlice := make([]string, 2, 5);
    newSlice[0] = "Filman";
    newSlice[1] = "Gulo";

    fmt.Println(newSlice);
    fmt.Println(len(newSlice));
    fmt.Println(cap(newSlice));

    lastSlice := make([]string, len(newSlice), cap(newSlice));
    copy(lastSlice, newSlice);

    fmt.Println(lastSlice);

//  Array VS Slice
    thisArray := [5]int{1,2,3,4,5,};
    thisSlice := []int{1,2,3,4,5,};

    fmt.Println(thisArray);
    fmt.Println(thisSlice);
}