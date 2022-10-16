package planetx

import (
	"fmt"
	"sort"
)

type Alien int

func (a Alien) String() string { return fmt.Sprintf("Alien %d", a) }

type Fight struct {
	City   string
	Alien1 Alien
	Alien2 Alien
}

func DistributeAliens(p Planet, rnd Randomizer, n int) map[Alien]string {
	if n < 0 {
		panic("cannot distribute negative amount of Aliens")
	}
	if n > len(p.cities) {
		panic("cannot distribute more aliens than cities")
	}
	aliens := make(map[Alien]string, n)

	cities := make([]string, 0, len(p.cities))
	for _, city := range p.cities {
		cities = append(cities, city.name)
	}
	rnd.Shuffle(len(p.cities), func(i, j int) {
		cities[i], cities[j] = cities[j], cities[i]
	})
	for i := 0; i < n; i++ {
		aliens[Alien(i+1)] = cities[i]
	}

	return aliens
}

func SimulateInvasion(
	p Planet,
	aliens map[Alien]string,
	rnd Randomizer,
	maxIteration int,
) ([]Fight, Planet) {
	fights := []Fight{}
	// copy aliens map
	// we should not modify the passed-in state
	aliens = func(m map[Alien]string) map[Alien]string {
		newm := make(map[Alien]string, len(m))
		for a, c := range m {
			newm[a] = c
		}
		return newm
	}(aliens)
	newPlanet := p.copyPlanet()

	// iterate over aliens in a deterministic manner
	alienNamesToIterate := func() (res []Alien) {
		for a := range aliens {
			res = append(res, a)
		}
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return
	}

	for i := 0; i < maxIteration; i++ {
		// move aliens to new locations
		movedLocations := make(map[string]Alien)
		for _, alien := range alienNamesToIterate() {
			alienCityName := aliens[alien]
			alienCity := newPlanet.cities[alienCityName]
			nearby := alienCity.nearbyCities()
			var visit *City
			if rnd.Int()%100 < 10 || len(nearby) == 0 { // TODO: extract these numbers out
				// should alien stay in the same place
				visit = alienCity
			} else {
				visit = nearby[rnd.Int()%len(nearby)]
			}

			if alienExistingHere, ok := movedLocations[visit.name]; ok {
				// this space is already occupied by some alien
				fights = append(fights, Fight{
					City:   visit.name,
					Alien1: alien,
					Alien2: alienExistingHere,
				})
				// removing aliens that just fought
				delete(aliens, alien)
				delete(aliens, alienExistingHere)
				delete(movedLocations, visit.name)
				visit.remove()
			} else {
				movedLocations[visit.name] = alien
			}

		}
		aliens = make(map[Alien]string)
		for loc, alien := range movedLocations {
			aliens[alien] = loc
		}
	}

	return fights, newPlanet
}