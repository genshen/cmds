package print

import (
	"flag"
	"fmt"
	"github.com/genshen/cmds"
)

var printCommand = &cmds.Command{
	Name:        "print",
	Summary:     "print some text",
	Description: "print some text, e.g. Hello World.",
	CustomFlags: false,
	HasOptions:  true,
}

var content string

func init() {
	printCommand.Runner = &version{}
	fs := flag.NewFlagSet("print", flag.ContinueOnError)
	printCommand.FlagSet = fs
	printCommand.FlagSet.StringVar(&content, "c", "default text", `text for print.`)
	printCommand.FlagSet.Usage = printCommand.Usage // use default usage provided by cmds.Command.
	cmds.AllCommands = append(cmds.AllCommands, printCommand)
}

type version struct{}

func (v *version) PreRun() error {
	return nil // if error != nil, function Run will be not execute.
}

func (v *version) Run() error {
	fmt.Println(content)
	return nil
}
