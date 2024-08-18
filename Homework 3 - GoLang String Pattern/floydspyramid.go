package main

import "fmt"

func main() {

    var i, j, rows int
    number := 1

    rows = 5

    fmt.Println("Floyd's Triangle")
    for i = 1; i <= rows; i++ {
        for j = 1; j <= i; j++ {
            fmt.Printf("%d ", number)
            number++
        }
        fmt.Println()
    }
}