package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	poem := []byte(`    Programming is fun
    When the work is done
    if you wanna make your work also fun:
    use Go!`)

	err := ioutil.WriteFile("poem.txt", poem, 0644)
	checkErr(err)

	data, err := ioutil.ReadFile("poem.txt")
	checkErr(err)

	fmt.Println(string(data))
}
