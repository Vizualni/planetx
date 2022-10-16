package planetx

import (
	"errors"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestUnmarshaling(t *testing.T) {
	for _, tt := range []struct {
		name      string
		in        string
		expPlanet Planet
		expErr    error
	}{
		{
			name:      "empty",
			in:        "",
			expPlanet: Planet{},
		},
		{
			name: "singly city",
			in:   "London",
			expPlanet: Planet{
				cities: map[string]*City{
					"London": {
						name: "London",
					},
				},
			},
		},
		{
			name: "two unconnected cities",
			in: `
			London
			New-York
			`,
			expPlanet: Planet{
				cities: map[string]*City{
					"London": {
						name: "London",
					},
					"New-York": {
						name: "New-York",
					},
				},
			},
		},
		{
			name: "one city defined, while other one is through road",
			in: `
			A north=B
			`,
			expPlanet: Planet{
				cities: map[string]*City{
					"A": {
						name: "A",
						north: &City{
							name: "B",
						},
					},
					"B": {
						name: "B",
					},
				},
			},
		},
		{
			name: "two cities defined with multiple directions",
			in: `
			A north=B south=C
			B east=A
			D north=E
			`,
			expPlanet: Planet{
				cities: map[string]*City{
					"A": {
						name: "A",
						north: &City{
							name: "B",
						},
						south: &City{
							name: "C",
						},
					},
					"B": {
						name: "B",
						east: &City{
							name: "A",
						},
					},
					"C": {
						name: "C",
					},
					"D": {
						name: "D",
						north: &City{
							name: "E",
						},
					},
					"E": {
						name: "E",
					},
				},
			},
		},
		{
			name: "invalid city name: starts with a number",
			in: `
			1A
			`,
			expErr: ErrCityNameInvalid,
		},
		{
			name: "invalid city name: contains a number",
			in: `
			aaaa1aaaa
			`,
			expErr: ErrCityNameInvalid,
		},
		{
			name: "invalid city name: contains symbol",
			in: `
			a.
			`,
			expErr: ErrCityNameInvalid,
		},
		{
			name: "invalid direction",
			in: `
			a down=b
			`,
			expErr: ErrInvalidDirection,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Unmarshal([]byte(tt.in))
			if !planetsEqual(got, tt.expPlanet) {
				t.Errorf("unmarshaled planet is not what was expected. got: %#v, expected: %#v", got, tt.expPlanet)
			}
			if !errors.Is(gotErr, tt.expErr) {
				t.Errorf("expected err: %v, got err: %v", tt.expErr, gotErr)
			}
		})
	}
}

func TestMarshaling(t *testing.T) {
	for _, tt := range []struct {
		name   string
		in     Planet
		expOut string
	}{
		{
			name:   "empty",
			in:     Planet{},
			expOut: "",
		},
		{
			name: "single city",
			in: Planet{
				cities: map[string]*City{
					"A": {
						name: "A",
					},
				},
			},
			expOut: "A",
		},
		{
			name: "two unconnected cities",
			in: Planet{
				cities: map[string]*City{
					"A": {
						name: "A",
					},
					"B": {
						name: "B",
					},
				},
			},
			expOut: `
			A
			B
			`,
		},
		{
			name: "two connected cities",
			in: Planet{
				cities: map[string]*City{
					"A": {
						name: "A",
					},
					"B": {
						name: "B",
						north: &City{
							name: "A",
						},
					},
				},
			},
			expOut: `
			A
			B north=A
			`,
		},
		{
			name: "connected cities and a stranded city",
			in: Planet{
				cities: map[string]*City{
					"A": {
						name: "A",
						north: &City{
							name: "B",
						},
					},
					"B": {
						name: "B",
					},
					"C": {
						name: "C",
					},
				},
			},
			expOut: `
			A north=B
			B
			C
			`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := string(Marshal(tt.in))

			// we need to bring the expected input and output to a same end
			// result so that we can properly compare them
			gotLinesRaw := strings.Split(got, "\n")
			gotLines := make([]string, 0, len(gotLinesRaw))

			expLinesRaw := strings.Split(tt.expOut, "\n")
			expLines := make([]string, 0, len(expLinesRaw))

			for _, line := range gotLinesRaw {
				line = strings.TrimSpace(line)
				if len(line) > 0 {
					gotLines = append(gotLines, line)
				}
			}

			for _, line := range expLinesRaw {
				line = strings.TrimSpace(line)
				if len(line) > 0 {
					expLines = append(expLines, line)
				}
			}

			sort.Strings(gotLines)
			sort.Strings(expLines)

			if !reflect.DeepEqual(gotLines, expLines) {
				t.Errorf("expected output: '%#v', got output: '%#v'", expLines, gotLines)
			}
		})
	}
}
