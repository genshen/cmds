package cmds

import (
	"fmt"
)

const DefaultUnknownSubCommandTemplate = `%s: unknown subcommand "%s"
Run '%s help' for usage.
`
// example: running './ProgramName UnknownSubCommand [options]', then this function will be executed.
var UnknownSubCommand = func(subCommand string) {
	name := GetProgramName()
	fmt.Printf(DefaultUnknownSubCommandTemplate,name, subCommand, name)
}

const DefaultUnknownSubCommandHelpTemplate = `Unknown help topic "%s". Run '%s help'.`
// example: running './ProgramName help UnknownSubCommand [options]', then this function will be executed.
var UnknownSubCommandHelp = func(subCommand string) {
	name := GetProgramName()
	fmt.Printf(DefaultUnknownSubCommandHelpTemplate, subCommand, name)
}
