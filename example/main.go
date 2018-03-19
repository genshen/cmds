package main
import (
	"github.com/genshen/cmds"
	_ "github.com/genshen/cmds/example/print"
	_ "github.com/genshen/cmds/example/version"
)
func main() {
	cmds.SetProgramName("example")
	cmds.Parse()
}
