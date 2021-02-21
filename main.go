package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	files := []string{
		"a", // TODO mnemonic info on file a
		// "b", // TODO mnemonic info on file b
		// "c", // TODO mnemonic info on file c
		// "d", // TODO mnemonic info on file d
		// "e", // TODO mnemonic info on file e
		// "f", // TODO mnemonic info on file f
	}

	for _, fileName := range files {
		fmt.Printf("****************** INPUT: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in", fileName))

		// TODO build input
		input := buildInput(inputSet)

		// TODO calculate and print input metrics
		printInputMetrics(input)

		// TODO execute algorithm
		result := algorithm(input)

		// TODO build output
		output := buildOutput(result)

		// TODO calculate and print result metrics
		printResultMetrics(result)

		output = strings.TrimSpace(output)
		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
	}
}
