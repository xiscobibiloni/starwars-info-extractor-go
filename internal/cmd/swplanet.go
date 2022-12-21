package cmd

import (
	"github.com/spf13/cobra"
	infoRetriever "github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
)

func InitSwPlanetsCmd(swapiClient infoRetriever.InfoRetriever) *cobra.Command {
	swpeople := &cobra.Command{
		Use:   "swplanets",
		Short: "Print swplanets",
		Run:   runCommandPlanetFn(swapiClient),
	}

	swpeople.Flags().StringP(fileFlag, shortFileFlag, "", "name of file to save response without extension")

	return swpeople
}
