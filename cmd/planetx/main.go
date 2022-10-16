package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/vizualni/planetx"
)

var (
	simulateInvasionFlags struct {
		Iterations int
		Aliens     int
		Seed       int64
		Path       string
	}
)

var (
	rootCmd = &cobra.Command{
		Use: "planetx",
	}

	simulateInvasionCmd = &cobra.Command{
		Use: "simulate-invasion",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := simulateInvasionFlags.Path
			if path == "" {
				return errors.New("must provide path")
			}

			var rnd planetx.Randomizer
			if simulateInvasionFlags.Seed == 0 {
				rnd = planetx.NewCryptoRandomize()
			} else {
				rnd = planetx.NewPseudoRandomizer(simulateInvasionFlags.Seed)
			}

			bz, err := ioutil.ReadFile(simulateInvasionFlags.Path)
			if err != nil {
				return fmt.Errorf("couldn't open file %s for reading: %w", path, err)
			}

			planet, err := planetx.Unmarshal(bz)
			if err != nil {
				return fmt.Errorf("could't unmarshal planetx file: %v", err)
			}

			aliens := planetx.DistributeAliens(planet, rnd, simulateInvasionFlags.Aliens)

			fights, newPlanet := planetx.SimulateInvasion(planet, aliens, rnd, simulateInvasionFlags.Iterations)

			for _, fight := range fights {
				fmt.Printf("City %s destroyed by aliens %s and %s fighting.\n", fight.City, fight.Alien1, fight.Alien2)
			}
			fmt.Println("========================")
			fmt.Println("A planet now looks like:")
			fmt.Println(string(planetx.Marshal(newPlanet)))

			return nil
		},
	}
)

func main() {
	simulateInvasionCmd.Flags().IntVar(&simulateInvasionFlags.Iterations, "iterations", 10000, "total number of iterations")
	simulateInvasionCmd.Flags().IntVar(&simulateInvasionFlags.Aliens, "aliens", 0, "number of aliens to send to a planet")
	simulateInvasionCmd.Flags().Int64Var(&simulateInvasionFlags.Seed, "seed", 0, "seed to use for the random behaviour. use 0 for cryptographic randomness")
	simulateInvasionCmd.Flags().StringVar(&simulateInvasionFlags.Path, "path", "", "path to the planetx file")

	rootCmd.AddCommand(simulateInvasionCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
