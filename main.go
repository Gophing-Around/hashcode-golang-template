package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Uncomment the line with the desired files (add other lines if needed)
	files := []string{ "a" }
	// files := []string{ "a", "b", "c", "d", "e", "f" }
	// files := []string{ "a", "b" }
	// files := []string{ "a", "b", "e", "f" }
	// files := []string{ "c" }
	// files := []string{ "d" }

	for _, fileName := range files {
		fmt.Printf("****************** INPUT: %s\n", fileName)
		inputSet := readFile(fmt.Sprintf("./inputFiles/%s.in", fileName))

		input := buildInput(inputSet)
		printInputMetrics(input)

		result := algorithm(input)

		output := buildOutput(result)
		printResultMetrics(result)
		
		ioutil.WriteFile(fmt.Sprintf("./result/%s.out", fileName), []byte(output), 0644)
	}
}
