package wrap

import (
	"fmt"
	"github.com/mkruse83/go-example-lib/lib"
)

func HelloWorld() {
	fmt.Println("Hello from wrapper v1.")
	lib.HelloWorld()
}
