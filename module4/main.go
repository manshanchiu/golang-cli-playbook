package module4

import (
	"fmt"

	"github.com/codemodus/kace"
)

func main() {
	fmt.Println("hello world")
	s := "this is a test sql."

	fmt.Println(kace.Camel(s))
}
