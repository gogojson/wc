package main

import (
	"fmt"
	"os"

	"github.com/gogojson/wc/cmd"
)

func main() {
	const wc = 342190
	const fileName = "/Users/jayson/Bucket/wc/text_maker/text.txt"

	cmd.Execute()

	// Read given file
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic("Cannot read file ")
	}
	fmt.Println(len(b))

}
