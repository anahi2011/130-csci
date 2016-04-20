package main
import "fmt"
func main() {
	for i := 0; i <= 100; i++ {
		if i%13 == 0 {
			fmt.Println(i, " -- hello")
		} else if i%2 == 0 {
			fmt.Println(i, " -- hola")
		} else if i%4 == 0 {
			fmt.Println(i, " -- something else")
		} else {
			fmt.Println(i)
		}
	}
}
