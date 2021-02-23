package num

import (
	"errors"
	"math"
	"math/rand"
)

func randIntRange(r *rand.Rand, min, max int) int {
	if min == max {
		return min
	}
	return r.Intn((max+1)-min) + min
}

func (f *Faker) Uint8() uint8 { return uint8Func(f.Rand) }

func uint8Func(r *rand.Rand) uint8 { return uint8(randIntRange(r, 0, math.MaxUint8)) }
