package planetx

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

// Randomizer is an interface that should return a random integer value and it can shuffle.
type Randomizer interface {
	Int() int
	Shuffle(n int, swap func(i, j int))
}

type cryptoRandSource struct{}

func (_ cryptoRandSource) Int63() int64 {
	var b [8]byte
	crand.Read(b[:])
	// mask off sign bit to ensure positive number
	return int64(binary.LittleEndian.Uint64(b[:]) & (1<<63 - 1))
}

func (cryptoRandSource) Seed(_ int64) {}

// NewPseudoRandomizer returns a pseudo randomizer that's seeded with an initial seed.
func NewPseudoRandomizer(seed int64) Randomizer {
	return rand.New(rand.NewSource(seed))
}

// NewCryptoRandomize returns a cryptographicaly safe random number generator.
func NewCryptoRandomize() Randomizer {
	return rand.New(cryptoRandSource{})
}
