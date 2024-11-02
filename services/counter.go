package services

import (
	"bufio"
	"fmt"
	"os"
)

func Counter(filename string, message string, split_func bufio.SplitFunc) (string, error) {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(split_func)

	counter := 0

	for scanner.Scan() {
		counter++
	}

	return fmt.Sprintf("%s %d", message, counter), nil
}
