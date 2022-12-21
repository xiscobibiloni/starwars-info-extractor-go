package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	inforetriever "github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
	"github.com/xiscobibiloni/starwars-info-extractor-go/internal/storage"
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

		fmt.Println("Save file CSV:")

		var csvStorage = storage.NewCsvStorage()
		csvStorage.StorageStarWarsPeople(people, file)

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

		fmt.Println("Save file CSV:")

		var csvStorage = storage.NewCsvStorage()
		csvStorage.StorageStarWarsPlanets(planets, file)
		return
	}
}
