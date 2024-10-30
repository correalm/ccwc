package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	SCAN_BYTES string = "bytes"
	SCAN_LINES        = "lines"
	SCAN_WORDS        = "words"
)

func main() {
	cli_values := map[string]string{
		"c": SCAN_BYTES,
		"l": SCAN_LINES,
		"w": SCAN_WORDS,
	}

	line_counter, byte_couter, word_counter := 0, 0, 0

	args := os.Args[1:]
	args_len := len(args)

	file_name := args[args_len-1]

	if args_len-1 == 0 {
		byte_couter = count_bytes(file_name)
		line_counter = count_lines(file_name)
		word_counter = count_words(file_name)
	} else {
		for _, arg := range args {
			switch action := cli_values[strings.Trim(arg, "-")]; action {
			case SCAN_BYTES:
				byte_couter = count_bytes(file_name)
			case SCAN_LINES:
				line_counter = count_lines(file_name)
			case SCAN_WORDS:
				word_counter = count_words(file_name)
			}
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
