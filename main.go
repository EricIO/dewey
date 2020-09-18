package main

import (
	"fmt"

	"github.com/EricIO/dewey/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
