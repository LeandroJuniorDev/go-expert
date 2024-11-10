package main

import (
	"bufio"
	"fmt"
	"os"
)
const fn= "file.txt"

func main() {

	// File writing
	cf, err := os.Create(fn)
	if err != nil {
		panic(err)
	}

	size, err := cf.Write([]byte("Hello, word!"))
	if err != nil {
		panic(err)
	}

	_ = cf.Close()

	fmt.Printf("File successfuly created: size = %d bytes\n", size)

	//File reading
	fc, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fc))

	//File reading part by part
	of, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(of)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
	_ = of.Close()

	//Detele file
	err = os.Remove(fn)
	if err != nil {
		panic(err)
	}
}