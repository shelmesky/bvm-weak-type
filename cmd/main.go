package main

import (
	"bvm/compiler"
	"bvm/parser"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
import "fmt"

func main() {
	inputFilename := os.Args[1]

	if len(inputFilename) == 0 {
		fmt.Println("need filename")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal("ReadFile failed:", err)
	}

	list := strings.Split(string(content), "\n")
	source := make([]string, 0, 32)

	for _, line := range list {
		source = append(source, line)
	}

	contractBody := strings.Join(source, "\r\n")

	root, err := parser.Parser(contractBody)
	if err != nil {
		panic(err)
	}

	err = compiler.Compile(root)

	fmt.Println(root)
}