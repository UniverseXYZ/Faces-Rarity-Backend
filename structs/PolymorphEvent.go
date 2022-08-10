package structs

import "math/big"

type PolymorphEvent struct {
	NewGene *big.Int
	OldGene *big.Int
	MorphId *big.Int
}
