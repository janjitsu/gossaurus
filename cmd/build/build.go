package main

import (
	"github.com/janjitsu/gossaurus/internal/command"
)

func main() {
	cmd := command.BuildWithArgs()
	cmd.Execute()
}
