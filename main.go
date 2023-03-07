package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func main() {
	args := os.Args[1:]

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fmt.Println("Read", n, "bytes:", string(data[:n]))

		for index, element := range repetition(string(data[:n])) {
			fmt.Println(index, "=", element)
		}
	}
	fmt.Println(args)
}
