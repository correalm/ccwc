### This is a implementation of [`wc Coding Challenge`](https://codingchallenges.fyi/challenges/challenge-wc)

### The goal is learn Go. The code is bad, for now.

### Build
`go build -o ccwc cmd/cli/main.go`

### Usage
  - Get all results (number of lines, words and bytes)
    `ccwc file_path`
  - Words
    `ccwc -w file_path`
  - Lines
    `ccwc -l file_path`
  - Bytes
    `ccwc -b file_path`
  - Mix
    `ccwc -l -w file_path`

### TODOS:
    - Testing
    - Maybe use goroutines to improve the performance
