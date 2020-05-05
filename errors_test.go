package cmds

import (
	"errors"
	"io"
	"testing"
)

func TestSubCommandParseError(t *testing.T) {
	err := SubCommandParseError{E: io.EOF}
	if !errors.Is(err, &SubCommandParseError{}) {
		t.Error("error type not match")
	}

	err2 := io.EOF
	if errors.Is(err2, &SubCommandParseError{}) {
		t.Error("error type not match")
	}
}
