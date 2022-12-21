package domain

type StorageInformation interface {
	StorageStarWarsPeople(people []StarWarsPeople, fileName string)
	StorageStarWarsPlanets(planets []StarWarsPlanet, fileName string)
}
