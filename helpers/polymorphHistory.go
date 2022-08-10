package helpers

import (
	"log"
	"rarity-backend/constants"
	"rarity-backend/metadata"
	"rarity-backend/structs"
)

// GetAttribute calcualtes the old and new attributes.
//
// This is later used to create a history snapshot of the polymorph
func GetAttribute(newGene string, oldGene string, geneIdx int, configService *structs.ConfigService) (structs.Attribute, structs.Attribute) {
	geneIdx = -geneIdx
	newAttribute := structs.Attribute{}
	oldAttribute := structs.Attribute{}

	if geneIdx > constants.BACKGROUND_GENE_START_IDX && geneIdx <= constants.BACKGROUND_GENE_END_IDX {
		newAttribute = metadata.GetBackgroundGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBackgroundGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.HAIR_LEFT_START_IDX && geneIdx <= constants.HAIR_LEFT_END_IDX {
		newAttribute = metadata.GetHairLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetHairLeftGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.HAIR_RIGHT_START_IDX && geneIdx <= constants.HAIR_RIGHT_END_IDX {
		newAttribute = metadata.GetHairRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetHairRightGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.EAR_LEFT_START_IDX && geneIdx <= constants.EAR_LEFT_END_IDX {
		newAttribute = metadata.GetEarsLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetEarsLeftGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.EAR_RIGHT_START_IDX && geneIdx <= constants.EAR_RIGHT_END_IDX {
		newAttribute = metadata.GetEarsRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetEarsRightGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.EYE_LEFT_START_IDX && geneIdx <= constants.EYE_LEFT_END_IDX {
		newAttribute = metadata.GetEyeLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetEyeLeftGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.EYE_RIGHT_START_IDX && geneIdx <= constants.EYE_RIGHT_END_IDX {
		newAttribute = metadata.GetEyeRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetEyeRightGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.BEARD_TOP_LEFT_START_IDX && geneIdx <= constants.BEARD_TOP_LEFT_END_IDX {
		newAttribute = metadata.GetBeardTopLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBeardTopLeftGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.BEARD_TOP_RIGHT_START_IDX && geneIdx <= constants.BEARD_TOP_RIGHT_END_IDX {
		newAttribute = metadata.GetBeardTopRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBeardTopRightGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.LIPS_LEFT_START_IDX && geneIdx <= constants.LIPS_LEFT_END_IDX {
		newAttribute = metadata.GetLipsLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetLipsLeftGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.LIPS_RIGHT_START_IDX && geneIdx <= constants.LIPS_RIGHT_END_IDX {
		newAttribute = metadata.GetLipsRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetLipsRightGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.BEARD_BOTTOM_LEFT_START_IDX && geneIdx <= constants.BEARD_BOTTOM_LEFT_END_IDX {
		newAttribute = metadata.GetBeardBottomLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBeardBottomLeftGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.BEARD_BOTTOM_RIGHT_START_IDX && geneIdx <= constants.BEARD_BOTTOM_RIGHT_END_IDX {
		newAttribute = metadata.GetBeardBottomRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBeardBottomRightGeneAttribute(oldGene, configService)
	} else {
		log.Println("Neshto se precaka da mu eba mamata ..")
	}

	return newAttribute, oldAttribute
}
