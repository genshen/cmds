package cmds

import (
	"flag"
	"fmt"
)

// AllCommands hold all available sub-commands.
// every Command instance of sub-command must be added to AllCommands in func init of sub-command.
// example of version.go:
/* var versionCommand = &cmds.Command{...}
  func init(){
    versionCommand.Runner = ... // an interface.
	cmds.AllCommands = append(cmds.AllCommands, versionCommand)
  }
*/
var AllCommands []*Command

// search all sub commands to find a command  by its unique name.
// If found, return the pointer of this command, return nil pointer otherwise.
func Find(name string) (bool, *Command) {
	for _, subCommand := range AllCommands {
		if subCommand.Name == name {
			return true, subCommand
		}
	}
	return false, nil
}

// interface for sub-commands.
// PreRun(such as checking necessary environment) execute before Run.
type CommandRunner interface {
	PreRun() error
	Run() error
}

// struct for every sub-command.
// every Command instance of sub-command must be added to AllCommands in func init of sub-command.
type Command struct {
	Name        string // sub-command name.
	Summary     string // summary of sub-command, it will be shown after running `./ProgramName help` for introducing every sub-command.
	Description string // description of sub-command, it will be shown after running `./ProgramName help sub-command` for introducing sub-command details.
	CustomFlags bool   // if true,it will parse args by itself.
	HasOptions  bool   // whether it has more options after sub-command.
	FlagSet     *flag.FlagSet
	Runner      CommandRunner
}

// print usage for sub-command.
func (c *Command) Usage() {
	fmt.Printf("%s\n\n", c.Description)
	fmt.Printf("Useage of command \"%s\":\n\n", c.Name)
	if c.HasOptions {
		fmt.Printf("\t%s %s [options]\n\n", GetProgramName(), c.Name) // todo more complex.
		fmt.Print("Options:\n\n")
		c.FlagSet.PrintDefaults()
	} else {
		fmt.Printf("\t%s %s\n\n", GetProgramName(), c.Name)
	}
}
