package main

import (
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.Execute()
}