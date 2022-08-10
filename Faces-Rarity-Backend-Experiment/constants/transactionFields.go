package constants

import "rarity-backend/structs"

var TxFieldNames = structs.TransactionFieldNames{
	ObjId:       "_id",
	BlockNumber: "blocknumber",
	TxIndex:     "txindex",
	TxHash:      "txhash",
	LogIndex:    "logindex",
}
