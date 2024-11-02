package services_test

import (
	"bufio"
	"ccwc/services"
	"fmt"
	"strings"
	"testing"
)

const FILE_NAME = "../__test__/fixtures/test.txt"

func TestCounterService(t *testing.T) {
  tests := []struct {
    name string
    message string
    filename string
    split_func bufio.SplitFunc
    result int
  } {
    {
      name: "returns the correct number of lines",
      message: "message 1",
      filename: FILE_NAME,
      split_func: bufio.ScanLines,
      result: 7145,
    },
    {
      name: "returns the correct number of bytes",
      message: "message 2",
      filename: FILE_NAME,
      split_func: bufio.ScanBytes,
      result: 342190,
    },
    {
      name: "returns the correct number of words",
      message: "message 3",
      filename: FILE_NAME,
      split_func: bufio.ScanWords,
      result: 58164,
    },
  }

  for _, test := range tests {
    t.Run(test.name, func(t *testing.T) {
      result, err := services.Counter(test.filename, test.message, test.split_func)
      expected := fmt.Sprintf(test.message, test.result)

      if err != nil {
        t.Errorf("An error occur on counter: %v", err)
      }

      if strings.Compare(result, expected) != 0 {
        t.Errorf("Expected %v; Actual %v", expected, result)
      }
    })
  }
}

func TestCounterServiceErrors(t *testing.T) {
  result, err := services.Counter("invalid_filename", "message", bufio.ScanBytes)

  if err == nil {
    t.Error("Invalid filename not trowns an error")
  }

  if strings.Compare(result, "") != 0 {
    t.Errorf("Expected empty string; Actual %v", result)
  }

  actualErrorMessage := err.Error()
  expectedErrorMessage := "open invalid_filename: no such file or directory"

  if strings.Compare(actualErrorMessage, expectedErrorMessage) != 0 {
    t.Errorf("Expected error message: %v; Actual: %v", expectedErrorMessage, actualErrorMessage)
  }
}
