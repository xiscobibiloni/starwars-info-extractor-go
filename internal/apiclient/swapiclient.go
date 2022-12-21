package apiclient

import (
	"encoding/json"
	"fmt"
	"github.com/xiscobibiloni/starwars-info-extractor-go/internal/domain"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	SwapiBaseUrl    = "https://swapi.dev/api"
	MethodGet       = "GET"
	PeopleOperation = "PEOPLE"
	PlanetOperation = "PLANET"
	UrlPeople       = "/people"
	UrlPlanet       = "/planets"
)

type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	HomeWorld string   `json:"homeworld"`
	Gender    string   `json:"gender"`
	Films     []string `json:"films"`
	Url       string   `json:"url"`
}

type ResponseStarWarsPeople struct {
	Count   int              `json:"count"`
	Results []StarWarsPeople `json:"results"`
}

type StarWarsPlanet struct {
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Terrain        string `json:"terrain"`
}

type ResponseStarWarsPlanet struct {
	Count   int              `json:"count"`
	Results []StarWarsPlanet `json:"results"`
}

type httpOperation struct {
	Path   string
	Method string
}

type clientConfiguration struct {
	baseUrl    string
	operations map[string]httpOperation
}

func NewSwapiClient() domain.InfoRetriever {
	return &clientConfiguration{
		baseUrl: SwapiBaseUrl,
		operations: map[string]httpOperation{
			PeopleOperation: {Path: UrlPeople, Method: MethodGet},
			PlanetOperation: {Path: UrlPlanet, Method: MethodGet},
		},
	}
}

var httpClient = &http.Client{
	Timeout: 20 * time.Second,
}

func (config *clientConfiguration) GetStarWarsPeople() (people []domain.StarWarsPeople, err error) {
	fmt.Printf("Base URL %s and path: %s\n\n", config.baseUrl, config.operations[PeopleOperation].Path)

	var starWarsPeople ResponseStarWarsPeople

	fmt.Println(fmt.Sprintf("%v%v", config.baseUrl, config.operations[PeopleOperation].Path))

	response, err := httpClient.Get(fmt.Sprintf("%v%v", config.baseUrl, config.operations[PeopleOperation].Path))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &starWarsPeople)
	if err != nil {
		return nil, err
	}

	var result []domain.StarWarsPeople
	for _, people := range starWarsPeople.Results {
		result = append(result, domain.StarWarsPeople{
			Name:      people.Name,
			Height:    people.Height,
			Mass:      people.Mass,
			HairColor: people.HairColor,
			HomeWorld: people.HomeWorld,
			Gender:    people.Gender,
			Films:     people.Films,
			Url:       people.Url,
		})
	}

	return result, nil
}

func (config *clientConfiguration) GetStarWarsPlanets() (planets []domain.StarWarsPlanet, err error) {
	fmt.Printf("Base URL %s and path: %s\n\n", config.baseUrl, config.operations[PeopleOperation].Path)

	var starWarsPlanets ResponseStarWarsPlanet

	fmt.Println(fmt.Sprintf("%v%v", config.baseUrl, config.operations[PlanetOperation].Path))

	response, err := httpClient.Get(fmt.Sprintf("%v%v", config.baseUrl, config.operations[PlanetOperation].Path))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &starWarsPlanets)
	if err != nil {
		return nil, err
	}

	var result []domain.StarWarsPlanet
	for _, planet := range starWarsPlanets.Results {
		result = append(result, domain.StarWarsPlanet{
			Name:           planet.Name,
			RotationPeriod: planet.RotationPeriod,
			OrbitalPeriod:  planet.OrbitalPeriod,
			Diameter:       planet.Diameter,
			Climate:        planet.Climate,
			Terrain:        planet.Terrain,
		})
	}

	return result, nil
}
