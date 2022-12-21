package domain

type StarWarsPeople struct {
	Name      string
	Height    string
	Mass      string
	HairColor string
	HomeWorld string
	Gender    string
	Films     []string
	Url       string
}

type StarWarsPlanet struct {
	Name           string
	RotationPeriod string
	OrbitalPeriod  string
	Diameter       string
	Climate        string
	Terrain        string
}
