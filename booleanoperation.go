package main

import "fmt"

func main() {
    var nilaiAkhir = 90;
    var kehadiran = 11;

    var lulusNilaiAkhir bool = nilaiAkhir > 80;
    var lulusKehadiran bool = kehadiran > 12;

    var lulus = lulusNilaiAkhir && lulusKehadiran;

    fmt.Println(lulusNilaiAkhir);
    fmt.Println(lulusKehadiran);
    fmt.Println(lulus);
}