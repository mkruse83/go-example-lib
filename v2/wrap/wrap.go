package wrap

import (
	"fmt"
	"github.com/mkruse83/go-example-lib/v2/lib"
)

func HelloWorld(_ string) {
	fmt.Println("Hello from wrapper v2.")
	lib.HelloWorld()
}
