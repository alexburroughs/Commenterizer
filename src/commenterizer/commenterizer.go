package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var LANGUAGE string

func main() {

	// get arguments excluding program
	args := os.Args[1:]

	LANGUAGE = args[1]
	codeString := readFile(args[0])

	for arr := codeString; len(arr) > 0; arr = arr[1:] {
		fmt.Println(arr[0])
	}

}

func readFile(filename string) []string {

	var inString []string

	// read data from filename
	data, err := ioutil.ReadFile(filename)

	// If error, exit
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	// get length of byte array
	//n := bytes.Index(data, []byte{0})

	// convert byte array to string
	s := string(data)

	// split string on newline
	inString = strings.Split(s, "\n")

	return inString
}

func commentLoop() {

}
