package cmd

import (
	"github.com/spf13/cobra"
	infoRetriever "github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
)

func InitSwPeopleCmd(swapiClient infoRetriever.InfoRetriever) *cobra.Command {
	swpeople := &cobra.Command{
		Use:   "swpeople",
		Short: "Print swpeople",
		Run:   runCommandPeopleFn(swapiClient),
	}

	swpeople.Flags().StringP(fileFlag, shortFileFlag, "", "name of file to save response without extension")

	return swpeople
}
