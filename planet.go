package planetx

import (
	"fmt"
	"reflect"
	"strings"
)

type zerosize struct{}

// City represents a city on a given planet x
type City struct {
	planet *Planet

	name string

	north *City
	south *City
	east  *City
	west  *City

	incoming map[*City]zerosize
}

// nearbyCities returns a list of cities where the current city is leading to
func (c *City) nearbyCities() []*City {
	cities := []*City{}
	if c.north != nil {
		cities = append(cities, c.north)
	}
	if c.south != nil {
		cities = append(cities, c.south)
	}
	if c.east != nil {
		cities = append(cities, c.east)
	}
	if c.west != nil {
		cities = append(cities, c.west)
	}
	return cities
}

// remove removes the city fromt the planet and removes all the connections in and out
func (c *City) remove() {
	for pointee := range c.incoming {
		if pointee.north == c {
			pointee.north = nil
		}
		if pointee.south == c {
			pointee.south = nil
		}
		if pointee.east == c {
			pointee.east = nil
		}
		if pointee.west == c {
			pointee.west = nil
		}
	}

	delete(c.planet.cities, c.name)
}

// addIncoming keeps a list of all cities that are pointing to a current one.
func (c *City) addIncoming(link *City) {
	if c.incoming == nil {
		c.incoming = make(map[*City]zerosize)
	}
	c.incoming[link] = zerosize{}
}

// Planet represent a planetX
type Planet struct {
	cities map[string]*City
}

// copyPlanet copies a planet. It's not optimized call, but rather is marshals
// the planet and then unmarshals it.
func (p Planet) copyPlanet() Planet {
	// TODO: implement a more performant copy of a planet
	// Note: a bit hackish way of copy-ing the planet ;)
	newp, err := Unmarshal(Marshal(p))
	if err != nil {
		panic(err)
	}
	return newp
}

// init sets the nil value planet
func (p *Planet) init() {
	if p == nil {
		p = &Planet{}
	}
	if p.cities == nil {
		p.cities = make(map[string]*City)
	}
}

// getCreateCity is a possible anti-pattern but it gets a city by name or it
// creates one by the provided name.
func (p *Planet) getCreateCity(name string) (*City, error) {
	err := verifyCityName(name)
	if err != nil {
		return nil, err
	}

	p.init()

	if city, ok := p.cities[name]; ok {
		return city, nil
	}
	city := &City{
		name:   name,
		planet: p,
	}
	p.cities[name] = city
	return city, nil
}

// planetsEqual checks if p1 and p2 planets are equal
func planetsEqual(p1, p2 Planet) bool {
	type cityMap map[string]map[string]string
	buildCityMap := func(p Planet) cityMap {
		cm := cityMap{}
		for _, c := range p.cities {
			directions := make(map[string]string)
			if c.north != nil {
				directions["north"] = c.north.name
			}
			if c.south != nil {
				directions["south"] = c.south.name
			}
			if c.east != nil {
				directions["east"] = c.east.name
			}
			if c.west != nil {
				directions["west"] = c.west.name
			}
			cm[c.name] = directions
		}
		return cm
	}

	return reflect.DeepEqual(buildCityMap(p1), buildCityMap(p2))
}

// addRoad adds a road between the two cities
func (from *City) addRoad(to *City, direction string) error {
	switch strings.ToLower(direction) {
	case "north":
		from.north = to
	case "south":
		from.south = to
	case "east":
		from.east = to
	case "west":
		from.west = to
	default:
		return fmt.Errorf("'%s': %w", direction, ErrInvalidDirection)
	}
	to.addIncoming(from)
	return nil
}
