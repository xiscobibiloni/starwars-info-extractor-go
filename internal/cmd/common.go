package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	inforetriever "github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
)

const (
	fileFlag      = "file"
	shortFileFlag = "f"
)

type CobraFn func(cmd *cobra.Command, args []string)

func runCommandPeopleFn(swapiClient inforetriever.InfoRetriever) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString(fileFlag)

		var people []inforetriever.StarWarsPeople
		people, _ = swapiClient.GetStarWarsPeople()

		fmt.Println(file)
		fmt.Println(people)

		return
	}
}
func runCommandPlanetFn(swapiClient inforetriever.InfoRetriever) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString(fileFlag)

		var planets []inforetriever.StarWarsPlanet
		planets, _ = swapiClient.GetStarWarsPlanets()

		fmt.Println(file)
		fmt.Println(planets)

		return
	}
}
