package structs

import (
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type RankMutex struct {
	Rank       int
	PrevRarity float64
	Operations []mongo.WriteModel
	Mutex      sync.Mutex
}
