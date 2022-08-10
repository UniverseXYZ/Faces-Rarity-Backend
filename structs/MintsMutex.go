package structs

import (
	"rarity-backend/models"
	"sync"
)

type MintsMutex struct {
	Mutex        sync.Mutex
	Mints        []models.PolymorphEntity
	TokensMap    map[string]bool
	Documents    []interface{}
	Transactions []interface{}
}
