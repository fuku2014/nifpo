package main

import (
	"fmt"
	"os"

	"github.com/fuku2014/nifpo/commands"
)

func main() {
	if err := commands.Nifpo.Execute(); err != nil {
		// TODO writerで書き出す
		fmt.Printf("nifcloud command is error: %s", err)
		os.Exit(1)
	}
}
