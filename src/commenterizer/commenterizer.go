package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Globals

// LANGUAGE the language of the code file being commented
var LANGUAGE string

// PARAMSTR the current string for parameters
var PARAMSTR string

// LANGMAP maps languages to function keywords
var LANGMAP map[string]string

// ParseState enum for parsing
type ParseState int

// Define Parsestate enumerator
const (
	Start      ParseState = 0
	Whitespace ParseState = 1
	InText     ParseState = 2
)

func main() {

	LANGMAP["go"] = "func"
	LANGMAP["bash"] = "function"

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

	// convert byte array to string
	s := string(data)

	// split string on newline
	inString = strings.Split(s, "\n")

	return inString
}

func commentLoop(codeString []string) string {

	var outString string

	for arr := codeString; len(arr) > 0; arr = arr[1:] {
		if isFunction(arr[0]) {

		}
	}

	// return the complete code file
	// with the comment headers added
	return outString
}

func getCommentHeader(funcHeader string, parameters []string) string {

	done := false
	header := ""

	// while the header is not done
	// (used if the user made a mistake)
	for !done {

	}

	return header
}

func isFunction(in string) bool {

	// get the line as bytes in order
	// to loop through the characters
	bytes := []byte(in)
	correct := true

	// loop through the line characters
	for i := 0; i < len(bytes); i++ {

		// get the language
		switch LANGUAGE {
		case "go":

			// set the function keyword
			sigStr := LANGMAP["go"]
			sig := []byte(sigStr)

			// if whitespace, ignore it
			if bytes[i] == ' ' {
				continue
			}

			// compare line start and function keyword
			for tmp := bytes[i:len(sig)]; len(tmp) > 0; {

				// If it is not a match then it is not a function
				if tmp[0] != sig[0] {
					return false
				}

				// loop through both strings
				tmp = tmp[1:]
				sig = sig[1:]
			}
		default:
			println(LANGUAGE + "is not supported")
			os.Exit(1)
		}
	}

	return correct
}

func getFunctionName(in string) string {

	// get the line as bytes in order
	// to loop through the characters
	bytes := []byte(in)

	// loop through the line characters
	for i := 0; i < len(bytes); i++ {
		// get the language
		switch LANGUAGE {
		case "go":

			// set the function keyword
			sigStr := LANGMAP["go"]
			sig := []byte(sigStr)

			// if whitespace, ignore it
			if bytes[i] == ' ' {
				continue
			}

			state := Start

			// compare line start and function keyword
			for tmp := bytes[i:len(sig)]; len(tmp) > 0; {

				switch state {

				case Start:
					// If it is not a match then it is not a function
					if tmp[0] != sig[0] {

					}
				case Whitespace:
				}

				// loop through both strings
				tmp = tmp[1:]
				sig = sig[1:]
			}
		default:
			println(LANGUAGE + "is not supported")
			os.Exit(1)
		}
	}
	return " "
}

/*func getFunctionParameters(in string) []string {

}*/
