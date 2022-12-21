package domain

type InfoRetriever interface {
	GetStarWarsPeople() ([]StarWarsPeople, error)
	GetStarWarsPlanets() ([]StarWarsPlanet, error)
}
