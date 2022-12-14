package apiclient

import (
	"fmt"
	infoRetriever "github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
)

const (
	MethodGet       = "GET"
	PeopleOperation = "PEOPLE"
	PlanetOperation = "PLANET"
	UrlPeople       = "/people"
	UrlPlanet       = "/planets"
)

type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    int      `json:"height"`
	Mass      int      `json:"mass"`
	HairColor string   `json:"hair_color"`
	HomeWorld string   `json:"homeworld"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
	Url       string   `json:"url"`
}

type StarWarsPlanet struct {
	Name           string `json:"name"`
	RotationPeriod int    `json:"rotation_period"`
	OrbitalPeriod  int    `json:"orbital_period"`
	Diameter       int    `json:"diameter"`
	Climate        string `json:"climate"`
	Terrain        string `json:"terrain"`
}

type httpOperation struct {
	Path   string
	Method string
}

type clientConfiguration struct {
	operations map[string]httpOperation
}

func NewSwapiClient() infoRetriever.InfoRetriever {
	return &clientConfiguration{map[string]httpOperation{
		PeopleOperation: {Path: UrlPeople, Method: MethodGet},
		PlanetOperation: {Path: UrlPlanet, Method: MethodGet},
	}}
}

func (config *clientConfiguration) GetStarWarsPeople() (people []infoRetriever.StarWarsPeople, err error) {
	fmt.Print(config.operations[PeopleOperation])
	return
}

func (config *clientConfiguration) GetStarWarsPlanets() (planets []infoRetriever.StarWarsPlanet, err error) {
	fmt.Print(config.operations[PlanetOperation])
	return
}
