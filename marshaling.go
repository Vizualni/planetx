package planetx

import (
	"fmt"
	"strings"
)

// Marshal takes the Planet and transforms it into:
// {city name} {direction1}={other city name} {directionN}={other city name}
func Marshal(p Planet) []byte {
	if len(p.cities) == 0 {
		return nil
	}
	var sb strings.Builder
	c := 0

	line := make([]string, 0, 5)
	for _, city := range p.cities {
		line = line[0:0:cap(line)]
		line = append(line, city.name)
		if city.north != nil {
			line = append(line, "north="+city.north.name)
		}
		if city.south != nil {
			line = append(line, "south="+city.south.name)
		}
		if city.west != nil {
			line = append(line, "west="+city.west.name)
		}
		if city.east != nil {
			line = append(line, "east="+city.east.name)
		}
		sb.WriteString(strings.Join(line, " "))
		c++
		if c < len(p.cities) {
			sb.WriteString("\n")
		}
	}

	// not really optimal to case from string to byte slice again
	return []byte(sb.String())
}

// Unmarshal takes the given data in the `{city name} {direction}={city name}`
// and builds the Planet.
func Unmarshal(data []byte) (Planet, error) {
	p := &Planet{
		cities: make(map[string]*City),
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// each line should be made out of these building blocks:
		// {city name} (north={city name})? (south={city name})? (east={city name})? (west={city name})?
		blocks := strings.Split(line, " ")
		if len(blocks) == 0 {
			continue
		}
		name, roads := blocks[0], blocks[1:]
		city, err := p.getCreateCity(name)
		if err != nil {
			return Planet{}, err
		}

		for _, road := range roads {
			roadBlock := strings.SplitN(road, "=", 2)
			if len(roadBlock) != 2 {
				return Planet{}, fmt.Errorf("line %s is invalid. road block '%s' ins not valid: %w", line, roadBlock, ErrInvalidInput)
			}
			direction, toCityName := roadBlock[0], roadBlock[1]
			toCity, err := p.getCreateCity(toCityName)
			if err != nil {
				return Planet{}, err
			}
			err = city.addRoad(toCity, direction)
			if err != nil {
				return Planet{}, err
			}
		}

		p.cities[name] = city
	}
	err := verifyPlanet(*p)
	if err != nil {
		return Planet{}, err
	}

	return *p, nil
}
