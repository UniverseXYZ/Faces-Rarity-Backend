package helpers

import (
	"rarity-backend/config"
	"rarity-backend/constants"
	"rarity-backend/metadata"
	"rarity-backend/models"
	"rarity-backend/structs"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

// CreateMorphEntity creates an entity which will be save in the rarities collection
func CreateMorphEntity(event structs.PolymorphEvent, metadata structs.Metadata, isVirgin bool, rarityResult structs.RarityResult) models.PolymorphEntity {
	var background, hairLeft, hairRight, earLeft, earRight, eyeLeft, eyeRight, beardTopLeft, beardTopRight, lipsLeft, lipsRight, beardBottomLeft, beardBottomRight structs.Attribute

	for _, attr := range metadata.Attributes {
		switch attr.TraitType {
		case constants.MorphAttriutes.Background:
			background = attr
		case constants.MorphAttriutes.HairLeft:
			hairLeft = attr
		case constants.MorphAttriutes.HairRight:
			hairRight = attr
		case constants.MorphAttriutes.EarLeft:
			earLeft = attr
		case constants.MorphAttriutes.EarRight:
			earRight = attr
		case constants.MorphAttriutes.EyeLeft:
			eyeLeft = attr
		case constants.MorphAttriutes.EyeRight:
			eyeRight = attr
		case constants.MorphAttriutes.BeardTopLeft:
			beardTopLeft = attr
		case constants.MorphAttriutes.BeardTopRight:
			beardTopRight = attr
		case constants.MorphAttriutes.LipsLeft:
			lipsLeft = attr
		case constants.MorphAttriutes.LipsRight:
			lipsRight = attr
		case constants.MorphAttriutes.BeardBottomLeft:
			beardBottomLeft = attr
		case constants.MorphAttriutes.BeardBottomRight:
			beardBottomRight = attr
		}
	}

	morphEntity := models.PolymorphEntity{
		TokenId:               int(event.MorphId.Int64()),
		Rank:                  0,
		CurrentGene:           event.NewGene.String(),
		HairLeft:              hairLeft.Value,
		HairRight:             hairRight.Value,
		EarLeft:               earLeft.Value,
		EarRight:              earRight.Value,
		EyeLeft:               eyeLeft.Value,
		EyeRight:              eyeRight.Value,
		BeardTopLeft:          beardTopLeft.Value,
		BeardTopRight:         beardTopRight.Value,
		LipsLeft:              lipsLeft.Value,
		LipsRight:             lipsRight.Value,
		BeardBottomLeft:       beardBottomLeft.Value,
		BeardBottomRight:      beardBottomRight.Value,
		Background:            background.Value,
		RarityScore:           rarityResult.ScaledRarity,
		IsVirgin:              isVirgin,
		ColorMismatches:       rarityResult.ColorMismatches,
		MainSetName:           rarityResult.MainSetName,
		MainMatchingTraits:    rarityResult.MainMatchingTraits,
		SecSetName:            rarityResult.SecSetName,
		SecMatchingTraits:     rarityResult.SecMatchingTraits,
		HasCompletedSet:       rarityResult.HasCompletedSet,
		HandsScaler:           rarityResult.HandsScaler,
		HandsSetName:          rarityResult.HandsSetName,
		MatchingHands:         rarityResult.MatchingHands,
		NoColorMismatchScaler: rarityResult.NoColorMismatchScaler,
		ColorMismatchScaler:   rarityResult.ColorMismatchScaler,
		VirginScaler:          rarityResult.VirginScaler,
		BaseRarity:            rarityResult.BaseRarity,
		ImageURL:              metadata.Image,
		Description:           metadata.Description,
		Name:                  metadata.Name,
	}
	if len(morphEntity.SecMatchingTraits) == 0 {
		morphEntity.SecMatchingTraits = []string{}
	}
	if len(morphEntity.MainMatchingTraits) == 0 {
		morphEntity.MainMatchingTraits = []string{}
	}
	return morphEntity
}

// CreateMorphSnapshot uses all the parameters in order to create a history snapshot of the polymorph. The morph cost mapping is updated depending on the morph type: Morph/Scramble.
//
//
// This snapshot is used to show the different variations each polymorph has gone through in the front end.
func CreateMorphSnapshot(geneDiff int, tokenId string, newGene string, oldGene string, timestamp uint64, oldAttr structs.Attribute, newAttr structs.Attribute, morphCostMap map[string]float32, configService *structs.ConfigService) models.PolymorphHistory {
	changeType, newAttrbiute, oldAttrubte := "", "", ""
	var newMorphCost float32 = 0
	morphCost := morphCostMap[tokenId]

	if morphCost == 0 {
		morphCost = config.SCRAMBLE_COST
	}

	if geneDiff <= 2 {
		changeType = "Morph"
		newAttrbiute = newAttr.Value
		oldAttrubte = oldAttr.Value
		newMorphCost = morphCost * 2
	} else {
		changeType = "Scramble"
		newAttrbiute = ""
		oldAttrubte = ""
		newMorphCost = config.SCRAMBLE_COST
	}
	morphCostMap[tokenId] = newMorphCost
	g := metadata.Genome(newGene)
	//character := metadata.GetBaseGeneAttribute(newGene, configService)
	genes := g.Genes()
	imageUrl := strings.Builder{}
	imageUrl.WriteString(constants.FACES_IMAGE_URL)

	for _, gene := range genes {
		imageUrl.WriteString(gene)
	}

	imageUrl.WriteString(".jpg")
	tokenInt, _ := strconv.Atoi(tokenId)

	return models.PolymorphHistory{
		TokenId:           tokenInt,
		Type:              changeType,
		DateTime:          time.Unix(int64(timestamp), 0).UTC(),
		AttributeChanged:  oldAttr.TraitType,
		PreviousAttribute: oldAttrubte,
		NewAttribute:      newAttrbiute,
		Price:             morphCost,
		ImageURL:          imageUrl.String(),
		NewGene:           newGene,
		OldGene:           oldGene,
	}
}

// SortMorphEvents sorts moprh events in chronological order(Block number -> Tx Index -> Log Index)
//
// If morph events aren't processed chronologically - the history snapshot of each morph event can have false information.
// We need the correct old state in order to calculate the correct differences in the gene
func SortMorphEvents(eventLogs []types.Log) {
	sort.Slice(eventLogs, func(i, j int) bool {
		curr := eventLogs[i]
		prev := eventLogs[j]

		if curr.BlockNumber < prev.BlockNumber {
			return true
		}

		if curr.BlockNumber > prev.BlockNumber {
			return false
		}

		if curr.TxIndex < prev.TxIndex {
			return true
		}

		if curr.TxIndex > prev.TxIndex {
			return false
		}

		if curr.Index < prev.Index {
			return true
		}

		if curr.Index > prev.Index {
			return false
		}
		return true
	})

}
