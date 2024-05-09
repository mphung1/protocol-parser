package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"project/parser"
	"strings"
)

func main() {
	inputValues := processArgs()

	fmt.Print("Processing input: ")
	for _, s := range inputValues {
		fmt.Print(s + " ")
	}
	fmt.Println()

	po := parser.ParserOutput{}

	err := parser.Parse(inputValues, &po)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range po.Output {
			fmt.Println(v)
		}
	}
}

func processArgs() (inputValues []string) {
	inputArgs := os.Args

	if len(inputArgs) == 1 {
		fmt.Println("Please specify the following")
		fmt.Println(" * Space separated list of input values")
		fmt.Println(" * The argument -f followed by the full path to an input file")
		os.Exit(0)
	} else if inputArgs[1] == "-f" {
		body, err := ioutil.ReadFile(inputArgs[2])
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		inputValues := strings.Split(string(body), " ")
		for i, val := range inputValues {
			inputValues[i] = strings.TrimSpace(val)
		}
		return inputValues
	} else {
		return inputArgs[1:]
	}
	return
}
