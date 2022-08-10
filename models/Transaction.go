package models

type Transaction struct {
	BlockNumber uint64 `json:"blocknumber,omitempty"`
	TxIndex     uint   `json:"txindex,omitempty"`
	TxHash      string `json:"txhash,omitempty"`
	LogIndex    uint   `json:"logindex,omitempty"`
}
