package wrap

import (
	"fmt"
	"github.com/mkruse83/go-example-lib/v3/lib"
)

func HelloWorld() {
	fmt.Println("Hello from wrapper v3.")
	lib.HelloWorld()
}
