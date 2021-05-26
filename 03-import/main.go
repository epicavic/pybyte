package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Print("The command line arguments are:")
	for i, v := range os.Args[:] {
		if i == 0 {
			v = filepath.Base(v)
		}
		fmt.Print(" ", v)
	}
	fmt.Println()
	fmt.Println("The GOROOT is:", runtime.GOROOT())
}
