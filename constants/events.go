package constants

import "rarity-backend/structs"

var MintEvent = structs.Event{
	Name:      "TokenMinted",
	Signature: "0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da",
}

var MorphEvent = structs.Event{
	Name:      "TokenMorphed",
	Signature: "0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b",
}

// 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef - TRANSFER EVENT
// 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925 - APPROVAL EVENT
// 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31 - APPROVAL FOR ALL EVENT
// 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d - ROLE GRANTED
