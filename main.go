package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	SCAN_BYTES string = "bytes"
	SCAN_LINES        = "lines"
	SCAN_WORDS        = "words"
)

func main() {
	cli_values := map[string]string{
		"-c": SCAN_BYTES,
		"-l": SCAN_LINES,
		"-w": SCAN_WORDS,
	}

	line_counter, byte_couter, word_counter := 0, 0, 0

	args := os.Args[1:]
	args_len := len(args)

	file_name := args[args_len-1]

	for count := 0; count <= args_len-1; count++ {
		action := cli_values[args[count]]

		execute_all := args_len-1 == 0

		if action == SCAN_BYTES || execute_all {
			byte_couter = count_bytes(file_name)
		}

		if action == SCAN_LINES || execute_all {
			line_counter = count_lines(file_name)
		}

		if action == SCAN_WORDS || execute_all {
			word_counter = count_words(file_name)
		}
	}

	fmt.Println("ARGS: ", line_counter, byte_couter, word_counter)
}

func count_lines(file_name string) int {
	file, err := os.Open(file_name)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		counter++
	}

	return counter
}

func count_bytes(file_name string) int {
	file, err := os.Open(file_name)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	counter := 0

	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		counter++
	}

	return counter
}

func count_words(file_name string) int {
	file, err := os.Open(file_name)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	counter := 0

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counter++
	}

	return counter
}
