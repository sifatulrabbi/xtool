package main

import "github.com/sifatulrabbi/xtool/internal/cli"

func main() {
	tool := cli.NewCli()
	if err := tool.TakeUserInput(); err != nil {
		panic(err)
	}
}
