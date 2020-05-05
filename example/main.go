package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/genshen/cmds"
	_ "github.com/genshen/cmds/example/print"
	_ "github.com/genshen/cmds/example/version"
)

func main() {
	cmds.SetProgramName("example")
	if err := cmds.Parse(); err != nil {
		if err == flag.ErrHelp {
		    return
		}
		// skip error in sub command parsing, because the error has been printed.
		if !errors.Is(err, &cmds.SubCommandParseError{}) {
			fmt.Println(err)
		}
	}
}
