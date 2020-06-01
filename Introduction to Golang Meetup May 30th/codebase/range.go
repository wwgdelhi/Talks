// _range_ iterates over elements in a variety of data
// structures. Let's see how to use `range` with some
// of the data structures we've already learned.

package main
import "fmt"

func main() {

	// Here we use `range` to sum the numbers in a slice.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` on arrays
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
}
