package planetx

import (
	"fmt"
	"regexp"
)

var (
	reCityName = regexp.MustCompile("^[a-zA-Z][a-zA-Z\\-]*$")
)

// verifyPlanet ensures that there is only a single occurance of a city
// and that city does not point to itself
func verifyPlanet(p Planet) error {
	visited := make(map[string]zerosize)
	for _, city := range p.cities {
		if _, ok := visited[city.name]; ok {
			return fmt.Errorf("city %s appeared twice in the graph: %w", city.name, ErrSameNodeMultiplied)
		}
		visited[city.name] = zerosize{}

		directionErr := func(direction string) error {
			return fmt.Errorf("city %s points to itself from %s: %w", city.name, direction, ErrCityPointsToItself)
		}

		// not checking using pointers because there might be a different node
		// but with a same name
		if city.east != nil {
			if city.east.name == city.name {
				return directionErr("east")
			}
		}
		if city.west != nil {
			if city.west.name == city.name {
				return directionErr("west")
			}
		}
		if city.north != nil {
			if city.north.name == city.name {
				return directionErr("north")
			}
		}
		if city.south != nil {
			if city.south.name == city.name {
				return directionErr("south")
			}
		}
	}
	return nil
}

func verifyCityName(name string) error {
	if !reCityName.MatchString(name) {
		return fmt.Errorf("'%s' is invalid. must start with a letter and can contain letters and dash only: %w", name, ErrCityNameInvalid)
	}

	return nil
}
