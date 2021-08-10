package wrap

import (
	"fmt"
	"github.com/mkruse83/go-example-lib/v2/lib"
)

func HelloWorld(_ int) {
	fmt.Println("Hello from wrapper v3.")
	lib.HelloWorld()
}
