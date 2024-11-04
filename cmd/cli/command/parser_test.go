package command_test

import (
	"bytes"
	"ccwc/cmd/cli/command"
	"ccwc/internal/interfaces"
	"os"
	"testing"
)

type Helper interface {
	Help()
}

type SpyHelper struct {
	Calls int
}

func (s *SpyHelper) Call() {
	s.Calls++
}

const FILENAME = "../../../__test__/fixtures/test.txt"

func TestParser(t *testing.T) {
	type fields struct {
		commands []interfaces.Command
	}
	type args struct {
		args []string
	}

	tests := []struct {
		name      string
		fields    fields
		args      args
		wantError bool
		helper    func()
		expected  string
	}{
		{
			name: "returns the number of lines when l arg is passed",
			fields: fields{
				commands: []interfaces.Command{
					command.NewLineCounter(FILENAME),
				},
			},
			args: args{
				args: []string{"-l"},
			},
			wantError: false,
			helper:    func() {},
			expected:  "Lines: 7145\n",
		},
		{
			name: "returns the number of words when w arg is passed",
			fields: fields{
				commands: []interfaces.Command{
					command.NewWordCounter(FILENAME),
				},
			},
			args: args{
				args: []string{"-w"},
			},
			wantError: false,
			helper:    func() {},
			expected:  "Words: 58164\n",
		},
		{
			name: "returns the number of bytes when b arg is passed",
			fields: fields{
				commands: []interfaces.Command{
					command.NewByteCounter(FILENAME),
				},
			},
			args: args{
				args: []string{"-b"},
			},
			wantError: false,
			helper:    func() {},
			expected:  "Bytes: 342190\n",
		},
		{
			name: "returns nothing when unknown arg is passed",
			fields: fields{
				commands: []interfaces.Command{
					command.NewByteCounter(FILENAME),
				},
			},
			args: args{
				args: []string{FILENAME, "-z"},
			},
			wantError: false,
			helper:    func() {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parser := command.NewParser(test.fields.commands)

			originalOut := os.Stdout

			r, w, _ := os.Pipe()
			os.Stdout = w

			err := parser.Parse(test.args.args, test.helper)

			w.Close()
			var buf bytes.Buffer
			buf.ReadFrom(r)

			os.Stdout = originalOut

			result := buf.String()

			if err != nil && !test.wantError {
				t.Errorf("parser.Parse() error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected: %v; Actual: %v", test.expected, result)
			}
		})
	}
}

func TestHelpCallWithNoArgs(t *testing.T) {
	parser := command.NewParser([]interfaces.Command{
		command.NewLineCounter(""),
	})

	spy := &SpyHelper{}

	err := parser.Parse([]string{}, spy.Call)

	if err != nil {
		t.Errorf("parser.Parse() error: %v", err)
	}

	if spy.Calls < 1 {
		t.Errorf("help function not called")
	}
}

func TestHelpCallWithArgs(t *testing.T) {
	parser := command.NewParser([]interfaces.Command{
		command.NewLineCounter(""),
	})

	spy := &SpyHelper{}

	err := parser.Parse([]string{"l"}, spy.Call)

	if err != nil {
		t.Errorf("parser.Parse() error: %v", err)
	}

	if spy.Calls > 0 {
		t.Errorf("help function called with args")
	}
}
