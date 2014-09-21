package main

import (
	"fmt"
	"github.com/coddingtonbear/go-jsonselect"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const VERSION string = "0.1"

var inputStream io.ReadCloser = os.Stdin

func main() {
	chars, err := ioutil.ReadAll(inputStream)

	selector := strings.Join(os.Args[1:], " ")
	parser, err := jsonselect.CreateParserFromString(string(chars))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(3)
	}

	results, err := parser.GetJsonElements(selector)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(4)
	}

	for _, result := range results {
		encoded, err := result.Encode()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Println(string(encoded))
	}
}
