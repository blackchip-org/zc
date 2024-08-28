package vars

import (
	"math/big"
	"math/rand/v2"

	"github.com/blackchip-org/zc/v6/pkg/coll"
)

const RandID = "rand"

type Rand struct {
	Rand   *rand.Rand
	Source *rand.PCG
	Seed   *big.Int
}

func ForRand(state coll.State) *Rand {
	v, ok := state.Var(RandID)
	if !ok {
		lo, hi := rand.Uint64(), rand.Uint64()
		src := rand.NewPCG(lo, hi)
		v = &Rand{
			Seed:   joinSeed(lo, hi),
			Source: src,
			Rand:   rand.New(src),
		}
		state.NewVar(RandID, v)
	}
	return v.(*Rand)
}

func (r *Rand) SplitSeed() (uint64, uint64) {
	var hi, lo, mask big.Int

	mask.SetString("ffffffffffffffff", 16)
	lo.Set(r.Seed)
	lo.And(&lo, &mask)

	hi.Set(r.Seed)
	hi.Rsh(&hi, 64)
	hi.And(&hi, &mask)

	return lo.Uint64(), hi.Uint64()
}

func joinSeed(lo uint64, hi uint64) *big.Int {
	var seed, l big.Int

	l.SetUint64(lo)

	seed.SetUint64(hi)
	seed.Lsh(&seed, 64)
	seed.Or(&seed, &l)
	return &seed
}
