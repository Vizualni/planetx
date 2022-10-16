package planetx

import (
	"fmt"
	"reflect"
	"strings"
)

type City struct {
	name string

	north *City
	south *City
	east  *City
	west  *City

	incoming map[*City]zerosize

	destroyed bool
}

func (c City) nearbyCities() []*City {
	// TODO: THIS IS WRONG!
	// need to have nodes which are pointing to this city
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

func (c *City) destroy() {
	// if anything here panics it means that the graph was incorrectly built!
	c.destroyed = true
}

func (c *City) addIncoming(link *City) {
	if c.incoming == nil {
		c.incoming = make(map[*City]zerosize)
	}
	c.incoming[link] = zerosize{}
}

type Planet struct {
	cities map[string]*City
}

func (p *Planet) init() {
	if p == nil {
		p = &Planet{}
	}
	if p.cities == nil {
		p.cities = make(map[string]*City)
	}
}

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
		name: name,
	}
	p.cities[name] = city
	return city, nil
}

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
