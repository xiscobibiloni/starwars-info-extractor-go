package storage

import (
	"encoding/csv"
	"github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
	"log"
	"os"
	"strings"
)

type storage struct {
}

func NewCsvStorage() domain.StorageInformation {
	return &storage{}
}

func (r *storage) StorageStarWarsPeople(people []domain.StarWarsPeople, fileName string) {
	var data [][]string
	for _, record := range people {
		row := []string{record.Name, record.Gender, record.Mass, record.HairColor, record.Height, record.HomeWorld, record.Url, strings.Join(record.Films, ",")}
		data = append(data, row)
	}
	csvFile, err := os.Create(fileName + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}

func (r *storage) StorageStarWarsPlanets(planets []domain.StarWarsPlanet, fileName string) {
	var data [][]string
	for _, record := range planets {
		row := []string{record.Name, record.Climate, record.Terrain, record.Diameter, record.OrbitalPeriod, record.RotationPeriod}
		data = append(data, row)
	}
	csvFile, err := os.Create(fileName + ".csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
