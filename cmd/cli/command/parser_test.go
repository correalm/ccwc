package command

import (
	"ccwc/internal/interfaces"
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
func TestHelpCallWithNoArgs(t *testing.T) {
  parser := &Parser{
    commands: []interfaces.Command{
      NewCountLines(""),
    },
  }

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
  parser := &Parser{
    commands: []interfaces.Command{
      NewCountLines(""),
    },
  }

  spy := &SpyHelper{}

  err := parser.Parse([]string{"l"}, spy.Call)

  if err != nil {
    t.Errorf("parser.Parse() error: %v", err)
  }

  if spy.Calls > 0 {
    t.Errorf("help function called with args")
  }
}
