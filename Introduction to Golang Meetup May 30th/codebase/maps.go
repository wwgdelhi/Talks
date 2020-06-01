// _Maps_ are Go's built-in [associative data type]
package main
import "fmt"

func main() {

	// To create an empty map, use the builtin `make`: `make(map[key-type]val-type)`.
	m := make(map[string]int)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. `fmt.Println` will show all of its key/value pairs.
	fmt.Println("map:", m)

	// Get a value for a key with `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin `delete` removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)
}
