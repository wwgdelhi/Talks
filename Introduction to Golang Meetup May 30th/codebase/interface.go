// _Interfaces_ are named collections of method signatures.
package main
import "fmt"
import "math"
// Here's a basic interface for geometric shapes.
type geometry interface {
    area() float64
}
// For our example we'll implement this interface on `circle` type.
type circle struct {
    radius float64
}
// To implement an interface in Go, we need to - implement all its methods.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
// If a variable has an interface type, then we can call methods that are in the named interface.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
}
func main() {
    c := circle{radius: 5}
    measure(c)
}
