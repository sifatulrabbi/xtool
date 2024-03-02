package cli

import "os"

/*
# CLI syntax
xtool [command] [params]
*/

type CLI struct {
	Command string
	Params  []CliParam
}

type CliParam struct {
	Name  string
	Value string
}

func NewCli() *CLI {
	cli := CLI{
		Command: "",
		Params:  []CliParam{},
	}
	return &cli
}

func (cli *CLI) TakeUserInput() error {
	args := os.Args[0:]
	return nil
}
