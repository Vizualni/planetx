package planetx

import (
	"testing"
)

func TestSimulation(t *testing.T) {
	for _, tt := range []struct {
		name      string
		planet    string
		expPlanet string
		aliensNum int
		seed      int64
	}{
		{
			name:      "empty",
			planet:    "",
			expPlanet: "",
			aliensNum: 0,
		},
		{
			name: "one city with one alien",
			planet: `
			A
			`,
			expPlanet: `
			A
			`,
			aliensNum: 1,
		},
		{
			name: "3 connected cities with two aliens",
			planet: `
			A north=B west=C
			B south=A
			C east=A
			`,
			expPlanet: `
			B
			C
			`,
			aliensNum: 2,
		},
		{
			name: "3 stranded cities with 3 aliens",
			planet: `
			A
			B
			C
			`,
			expPlanet: `
			A
			B
			C
			`,
			aliensNum: 3,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			expPlanet, err := Unmarshal([]byte(tt.expPlanet))
			if err != nil {
				panic(err)
			}
			rnd := NewPseudoRandomizer(tt.seed)
			planet, err := Unmarshal([]byte(tt.planet))
			if err != nil {
				panic(err)
			}
			aliens := DistributeAliens(planet, rnd, tt.aliensNum)
			_, newPlanet := SimulateInvasion(planet, aliens, rnd, 10000)
			if !planetsEqual(expPlanet, newPlanet) {
				t.Errorf("planets are not equal. expected: %#v, got: %#v", expPlanet, newPlanet)
			}

		})
	}
}
