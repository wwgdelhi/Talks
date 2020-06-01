package main
import "fmt"

func main() { 
    i := 1
    for i <= 3 {  // The most basic type, with a single condition.
        fmt.Println(i)
        i = i + 1
    }

    // `for` without a condition will loop repeatedly until you `break` out of it
    for {
        fmt.Println("loop")
        break
    }
    
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue // You can also `continue` to the next iteration of the loop.
        } 
        fmt.Println(n)
    }
}