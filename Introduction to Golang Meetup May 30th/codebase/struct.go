// Go's _structs_ are typed collections of fields.
package main
import "fmt"
type person struct { // This `person` struct type has `name` and `age` fields.
    name string
    age  int
}
func main() {
    // This syntax creates a new struct.
    fmt.Println(person{"Bob", 20})

    // You can name the fields when initializing a struct.
    fmt.Println(person{name: "Alice", age: 30})

    // Omitted fields will be zero-valued.
    fmt.Println(person{name: "Fred"})

    // Access struct fields with a dot.
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // Structs are mutable.
    s.age = 51
    fmt.Println(s.age) }
