// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    // This call will be made synchronously    
    f("direct")

    // Execute asynchronously
    go f("goroutine")

    // Start an anonymous gorouting
    go func(msg string) {
        fmt.Println(msg)
    }("anonymous")

    // Blocking call, waiting for input from user
    // If this not here, possible never see the output of our
    // two gouroutines
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}