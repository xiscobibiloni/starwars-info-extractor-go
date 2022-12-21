package main

import (
	"github.com/spf13/cobra"
	"github.com/xiscobibiloni/starwars-info-extractor-go/internal/apiclient"
	"github.com/xiscobibiloni/starwars-info-extractor-go/internal/cmd"
)

func main() {

	var swapiClient = apiclient.NewSwapiClient()

	rootCmd := &cobra.Command{Use: "starwars-cli"}
	rootCmd.AddCommand(cmd.InitSwPeopleCmd(swapiClient))
	rootCmd.AddCommand(cmd.InitSwPlanetsCmd(swapiClient))
	rootCmd.Execute()
}
