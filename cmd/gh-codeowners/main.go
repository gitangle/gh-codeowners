package main

import (
	"os"

	"github.com/gitangle/gh-codeowners/cmd/gh-codeowners/cmd"
)

func main() {
	rootCmd := cmd.NewRoot()
	if err := rootCmd.Execute(); err != nil {
		// error is already printed by cobra, we can here add error switch
		// in case we would like to exit with different codes
		os.Exit(1)
	}
}
