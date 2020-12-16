package main

import (
	"os"

	"github.com/bloxapp/pools-network/cmd/pools-networkd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
