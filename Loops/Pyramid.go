package main

import "fmt"

func main(){

    var n int
    fmt.Printf (" Enter n value: \n")
    fmt.Scanf("%d", &n)

    for  i :=1; i <= n; i++{

        for j := 1; j <= n - i; j++{
            fmt.Printf ("  ")
        }

        for  k := 1; k <= ( 2 * i - 1); k++{
            fmt.Printf ("* ")
        }
       fmt.Printf ("\n")
    }
}
