package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)

	for i, arg := range os.Args {
		if i > 0 {
			printFile(arg)
		}
	}
}

func printFile(s string) {
	fmt.Println("\nfilename: ", s)
	file, err := os.Open(s) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
}
