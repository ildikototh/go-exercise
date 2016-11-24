package main //the package this file contains
import "fmt" //import other packages
func add(a,b int) int {
    return a+b
}

func main() {
    c := add(10,10)
	fmt.Println(c)
}
