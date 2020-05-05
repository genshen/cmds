package cmds

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var programName string

func SetProgramName(name string) {
	programName = name
}

// return value depends on the var programName, if programName is empty,return os.Args[0], otherwise it will return var programName.
func GetProgramName() string {
	if programName == "" {
		return os.Args[0]
	} else {
		return programName
	}
}

const DefaultOptionsFormat = `` // for printing subcommand option.
// default usage function.
var Usage = func() {
	name := GetProgramName()
	fmt.Printf("Usage of %s:\n\n", name)
	fmt.Printf("\t%s command [arguments]\n\n", name)
	fmt.Print("The commands are:\n\n")
	//PrintDefaultOptions()
	for _, c := range AllCommands { // print sub-command option.
		fmt.Printf("\t%s\t\t%s\n", c.Name, c.Summary)
	}
	// fmt.Printf("\t%s\t\t%s\n", "help", "show this help")

	fmt.Println()
	fmt.Printf(`Use "%s help [command]" for more information about a command.`, name)
	fmt.Println()
}

// it will be called when args have only a program name.
var SingleArg = func() {
	Usage()
}

// parse args from os.Args.
// and dispatch them to sub-command handle in AllCommands.
func Parse() error {
	if len(os.Args) <= 1 { // if args have only a program name.
		SingleArg() // the default method is to print usage
		return nil
	}

	args := os.Args[1:]
	// 'help' command
	if args[0] == "help" || args[0] == "--help" || args[0] == "h" || args[0] == "-h" {
		Help(args[1:])
		return flag.ErrHelp
	}

	// find a available subCommand, and pass all left args (except command name) to this subCommand.
	for _, subCommand := range AllCommands {
		if subCommand.Name == args[0] {
			if !subCommand.CustomFlags { // otherwise, handle parse by sub-command itself
				err := subCommand.FlagSet.Parse(args[1:])
				// the ErrorHandling here must be flag.ContinueOnError
				// other types error handling has exit program
				if err == flag.ErrHelp {
					// does not return error if it is `--help`
					return err
				} else if err != nil {
					// error message and usage have been printed in FlagSet.Parse.
					return fmt.Errorf("%w", SubCommandParseError{E: err})
				} // otherwise, continue run
			}
			if subCommand.Runner != nil {
				if err := subCommand.Runner.PreRun(); err == nil {
					return subCommand.Runner.Run()
					// todo rollback if error
				} else {
					return err
				}
			}
			// todo error output.
			return errors.New("the sub-command does not implement the `Runner` interface")
		}
	}
	return UnknownSubCommand(args[0])
}

// print help messages, including 'programName --help' and 'programName --help sub-command'.
// args: os.Args[2:]
func Help(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}

	// have more help flag.
	for _, c := range AllCommands {
		if c.Name == args[0] {
			c.FlagSet.Usage()
			return
		}
	}
	UnknownSubCommandHelp(args[0]) // print message of unkonwn sub-command.
}
