package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Format("2006-01-02 15:04:05.000000")
	s := fmt.Sprintf("hello world at %s", t)
	fmt.Println(s)
}
