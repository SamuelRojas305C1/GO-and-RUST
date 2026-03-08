Package main
import "fmt"

func sumar(a int, b int) int {
	return a + b
}
func main() {
	resultado := sumar(10, 20)
	fmt.Prinln("La suma es:", resultado)
}