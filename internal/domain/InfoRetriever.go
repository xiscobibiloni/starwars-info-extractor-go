package domain

type StarWarsPeople struct {
	Name      string
	Height    int
	Mass      int
	HairColor string
	HomeWorld string
	Gender    string
	Films     []string
	Url       string
}

type StarWarsPlanet struct {
	Name           string
	RotationPeriod int
	OrbitalPeriod  int
	Diameter       int
	Climate        string
	Terrain        string
}

type InfoRetriever interface {
	GetStarWarsPeople() ([]StarWarsPeople, error)
	GetStarWarsPlanets() ([]StarWarsPlanet, error)
}
