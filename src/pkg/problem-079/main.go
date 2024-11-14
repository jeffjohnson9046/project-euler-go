package main

import (
	"fmt"
	"projecteuler/src/pkg/diagnostics"
	"projecteuler/src/pkg/fileutils"
	"strings"
	"time"
)

// Read the provided file.
// func readFile(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	defer func() {
// 		if err := file.Close(); err != nil {
// 			log.Fatalf("Error reading file %s: %v", path, err)
// 		}
// 	}()

// 	fileScanner := bufio.NewScanner(file)
// 	fileScanner.Split(bufio.ScanLines)

// 	var fileLines []string
// 	for fileScanner.Scan() {
// 		fileLines = append(fileLines, fileScanner.Text())
// 	}

// 	return fileLines, nil
// }

// Sort the graph of digits
func topologicalSort(graph map[rune][]rune) string {
	inDegree := make(map[rune]int)

	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}

		if _, ok := inDegree[node]; !ok {
			inDegree[node] = 0
		}
	}

	var result strings.Builder
	var queue []rune

	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result.WriteRune(node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--

			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result.String()) != len(inDegree) {
		return "" // Cycle detected; this shouldn't happen given the data and the rules of this problem
	}

	return result.String()
}

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "problem-079")

	path := "./resources/login-attempts.txt"
	loginAttempts, err := fileutils.ReadFile(path)
	if err != nil {
		fmt.Printf("Error attempting to read login attempst file %s: %v", path, err)
	}

	digitsGraph := make(map[rune][]rune)

	for _, attempt := range loginAttempts {
		for i := 0; i < len(attempt)-1; i++ {
			from, to := rune(attempt[i]), rune(attempt[i+1])

			digitsGraph[from] = append(digitsGraph[from], to)
		}
	}

	for key, value := range digitsGraph {
		fmt.Printf("key: %s, value: %s\n", string(key), string(value))
	}

	result := topologicalSort(digitsGraph)
	if result == "" {
		fmt.Println("\nBONED")
	} else {
		fmt.Printf("the passcode is: %s\n", result)
	}
}
