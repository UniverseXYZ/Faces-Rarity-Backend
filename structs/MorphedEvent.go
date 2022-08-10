package structs

import "math/big"

type MorphedEvent struct {
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType uint8
}
