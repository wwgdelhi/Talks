package main
import "fmt"

func fac(n int) int {
	if n == 0{
		return 1
	}
	return fac(n-1) * n;
}
func main(){
	fmt.Println(fac(10));
}