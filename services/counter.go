package services

import (
	"bufio"
	"fmt"
	"os"
)

func Counter(filename string, message string, split_func bufio.SplitFunc) error {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return fmt.Errorf("Error when open file %s: %v", filename, err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(split_func)

	counter := 0

	for scanner.Scan() {
		counter++
	}
	fmt.Println(message, counter)

	return nil
}
