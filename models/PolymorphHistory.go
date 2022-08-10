package models

import "time"

type PolymorphHistory struct {
	Type              string    `json:"type,omitempty"`
	TokenId           int       `json:"tokenid,omitempty"`
	DateTime          time.Time `json:"datetime,omitempty"`
	AttributeChanged  string    `json:"attributechanged,omitempty"`
	PreviousAttribute string    `json:"previousattribute,omitempty"`
	NewAttribute      string    `json:"newattribute,omitempty"`
	Price             float32   `json:"price,omitempty"`
	ImageURL          string    `json:"imageurl,omitempty"`
	NewGene           string    `json:"newgene,omitempty"`
	OldGene           string    `json:"oldgene,omitempty"`
	Character         string    `json:"character,omitempty"`
}
