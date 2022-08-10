package services

import (
	"fmt"
	"log"
	"math"
	"rarity-backend/config"
	"rarity-backend/constants"
	"rarity-backend/helpers"
	"rarity-backend/structs"
	"strings"
)

// CalulateRarityScore is the core function responsible for calcualting the rarity score.
//
// It calculates the rarity score of the polymorph, the different scalers used in the formuala and other rarity related metadata that is tracked and stored in the database.
//
// Configurations can be found in rarityConfig.go
func CalulateRarityScore(attributes []structs.Attribute, isVirgin bool) structs.RarityResult {
	leftHand, rightHand, _ := parseAttributes(attributes)

	hasCompletedSet, mainSetName, mainMatchingTraits, secSetname, secMatchingTraits := calculateCompleteSets(attributes)
	mainSetName, mainMatchingTraits, secSetname, secMatchingTraits, isColoredSet, colorMismatches := getColorMismatches(attributes, mainSetName, mainMatchingTraits, secSetname, secMatchingTraits)
	scalers := getScalers(hasCompletedSet, mainSetName, colorMismatches, isVirgin, isColoredSet)
	handsScaler, handsSetName, matchingHandsCount, mainMatchingTraitsWithHands := getFullSetHandsScaler(mainMatchingTraits, hasCompletedSet, mainSetName, leftHand, rightHand)

	mainSetCount := float64(len(mainMatchingTraits))
	secSetBonus := config.SECONDARY_SET_SCALER * float64(len(secMatchingTraits))
	mismatchPenalty := config.MISMATCH_PENALTY * colorMismatches

	baseRarity := math.Pow(2, mainSetCount-mismatchPenalty+secSetBonus)

	totalScalars := scalers.NoColorMismatchScaler * handsScaler * scalers.VirginScaler
	scaledRarity := math.Round(((baseRarity * totalScalars * 100) * 100)) / 100
	log.Println("Rarity index: " + fmt.Sprintf("%f", (scaledRarity)))

	return structs.RarityResult{
		HasCompletedSet:       hasCompletedSet,
		MainSetName:           mainSetName,
		MainMatchingTraits:    mainMatchingTraitsWithHands,
		SecSetName:            secSetname,
		SecMatchingTraits:     secMatchingTraits,
		ColorMismatches:       int(colorMismatches),
		HandsScaler:           handsScaler,
		HandsSetName:          handsSetName,
		MatchingHands:         matchingHandsCount,
		NoColorMismatchScaler: scalers.NoColorMismatchScaler,
		ColorMismatchScaler:   scalers.ColorMismatchScaler,
		VirginScaler:          scalers.VirginScaler,
		BaseRarity:            baseRarity,
		ScaledRarity:          scaledRarity,
	}
}

//parseAttributes parses the array of attrbutes.
//
//Returns left hand, right hand and the attrbitues without Character, Background attrbiutes as they aren't used in the rarity score formula
func parseAttributes(attributes []structs.Attribute) (structs.Attribute, structs.Attribute, []structs.Attribute) {
	var hairLeft, hairRight structs.Attribute
	var rarityAttributes []structs.Attribute

	for _, attr := range attributes {
		switch attr.TraitType {
		case constants.MorphAttriutes.Background:
		case constants.MorphAttriutes.HairLeft:
			hairLeft = attr
		case constants.MorphAttriutes.HairRight:
			hairRight = attr
		default:
			rarityAttributes = append(rarityAttributes, attr)
		}
	}

	return hairLeft, hairRight, rarityAttributes
}

// getScalers calculates the eligible scalers for the polymorph
func getScalers(hasCompletedSet bool, setName string, colorMismatches float64, isVirgin bool, isColoredSet bool) structs.Scalers {
	var noColorMismatchScaler, colorMismatchScaler, virginScaler float64 = 1, 1, 1

	if hasCompletedSet && isColoredSet && colorMismatches == 0 {
		noColorMismatchScaler = config.NO_COLOR_MISMATCH_SCALER
	} else if hasCompletedSet && isColoredSet && colorMismatches != 0 {
		colorMismatchScaler = config.COLOR_MISMATCH_SCALER
	}

	if isVirgin {
		virginScaler = config.VIRGIN_SCALER
	}

	return structs.Scalers{
		ColorMismatchScaler:   colorMismatchScaler,
		NoColorMismatchScaler: noColorMismatchScaler,
		VirginScaler:          virginScaler,
	}
}

// getColorMismatches calculates determines if the set has colors or not and the number of color mismatches if applicable.
//
// Also swaps main with secondary set if both are colored, of equal length and sec set has less color mismatches
//
// Color sets can be found in rarityConfig.go
func getColorMismatches(attributes []structs.Attribute, mainSetName string, mainMatchingTraits []string, secSetName string, secMatchingTraits []string) (string, []string, string, []string, bool, float64) {
	// Check if main set is with colors
	colorSetMain, colorSetSecondary := structs.ColorSet{}, structs.ColorSet{}
	if strings.Contains(mainSetName, config.FootbalSetWithColors.Name) {
		colorSetMain = config.FootbalSetWithColors
	} else if strings.Contains(mainSetName, config.SpartanSetWithColors.Name) {
		colorSetMain = config.SpartanSetWithColors
	} else if strings.Contains(mainSetName, config.KnightSetWithColors.Name) {
		colorSetMain = config.KnightSetWithColors
	} else {
		// Set is without colors
		return mainSetName, mainMatchingTraits, secSetName, secMatchingTraits, false, 0
	}

	// If main set has colors we check if secondary set also has colors. We may need to replace main with secondary if both are of
	// equal length and secondary has less color mismatches
	if strings.Contains(secSetName, config.FootbalSetWithColors.Name) {
		colorSetSecondary = config.FootbalSetWithColors
	} else if strings.Contains(secSetName, config.SpartanSetWithColors.Name) {
		colorSetSecondary = config.SpartanSetWithColors
	} else if strings.Contains(secSetName, config.KnightSetWithColors.Name) {
		colorSetSecondary = config.KnightSetWithColors
	}

	mainSetColorMismatches := getColorSetMismatches(attributes, colorSetMain, mainSetName)
	if len(mainMatchingTraits) == len(secMatchingTraits) && colorSetMain.Name != "" && colorSetSecondary.Name != "" {
		secSetColorMismatches := getColorSetMismatches(attributes, colorSetSecondary, secSetName)
		if secSetColorMismatches < mainSetColorMismatches {
			mainMatchingTraits, secMatchingTraits = secMatchingTraits, mainMatchingTraits
			mainSetName, secSetName = secSetName, mainSetName
			mainSetColorMismatches = secSetColorMismatches
		}
	}

	return mainSetName, mainMatchingTraits, secSetName, secMatchingTraits, true, mainSetColorMismatches
}

//getColorSetMismatches returns the color mismatches of a colored set
func getColorSetMismatches(attributes []structs.Attribute, coloredSet structs.ColorSet, setName string) float64 {
	colorMap := make(map[string]float64)
	var totalColorsOccurances, primaryColorOccurances float64
	var primaryColor string

	for _, attr := range attributes {
		for _, color := range coloredSet.Colors {
			if strings.Contains(attr.Value, color) && helpers.StringInSlice(setName, attr.Sets) {
				totalColorsOccurances++
				colorMap[color]++
				break
			}
		}
	}

	for k, v := range colorMap {
		if primaryColorOccurances < v {
			primaryColorOccurances = v
			primaryColor = k
		}
	}

	colorMismatches := totalColorsOccurances - primaryColorOccurances

	if coloredSet.Name == "Spartan" || coloredSet.Name == "Knight" {
		if primaryColor == "Silver" || primaryColor == "Golden" {
			// Silver and Golden Spartan sets have Brown Footwear. That's why we convert the brown to silver in order to detect correct number of color mismatches
			if colorMap["Brown"] != 0 {
				colorMismatches--
			}
		}
	} else if coloredSet.Name == "Football Star" {
		// We have to manually iterate over the attributes and count each attribute towards a set to be able to count the color mismatches
		awayAttrCount, homeAttrCount := 0, 0
		for _, attr := range attributes {
			switch attr.Value {
			case "Red Football Helmet", "White Football Jersey", "Red Football Pants", "White/Yellow Football Cleats":
				awayAttrCount++
			case "Grey Football Helmet", "Red Football Jersey", "Grey Football Pants", "Red Football Cleats":
				homeAttrCount++
			}
		}
		if homeAttrCount == 0 || awayAttrCount == 0 {
			colorMismatches = 0
		} else if homeAttrCount == awayAttrCount {
			colorMismatches = float64(homeAttrCount)
		} else if homeAttrCount > awayAttrCount {
			colorMismatches = float64(homeAttrCount - awayAttrCount)
		} else {
			colorMismatches = float64(awayAttrCount - homeAttrCount)
		}
	}
	return colorMismatches
}

// getFullSetHandsScaler calculates the correct hands scaler based on the state of the set(no, incomplete or completed set)
func getFullSetHandsScaler(mainMatchingTraits []string, hasCompletedSet bool, completedSetName string,
	leftHandAttr structs.Attribute, rightHandAttr structs.Attribute) (float64, string, int, []string) {
	var matchingSetHandsCount int

	// Match left hand
	for _, set := range leftHandAttr.Sets {
		if set == completedSetName {
			matchingSetHandsCount++
			// mainMatchingTraits = append(mainMatchingTraits, leftHandAttr.TraitType)
			break
		}
	}

	// Match right hand
	for _, set := range rightHandAttr.Sets {
		if set == completedSetName {
			matchingSetHandsCount++
			// mainMatchingTraits = append(mainMatchingTraits, leftHandAttr.TraitType)
			break
		}
	}

	if matchingSetHandsCount == 0 {
		// Check if they are matching set
		handMap := make(map[string]int)
		for _, set := range leftHandAttr.Sets {
			handMap[set]++
		}
		for _, set := range rightHandAttr.Sets {
			handMap[set]++
			if handMap[set] == 2 {
				if leftHandAttr.Value == rightHandAttr.Value {
					return config.NO_SET_TWO_SAME_MATCHING_HANDS_SCALER, set, handMap[set], mainMatchingTraits
				}
			}
		}
	} else if len(mainMatchingTraits) == 2 &&
		(mainMatchingTraits[0] == "Right-Hand" || mainMatchingTraits[0] == "Left-Hand") &&
		(mainMatchingTraits[1] == "Right-Hand" || mainMatchingTraits[1] == "Left-Hand") {
		if leftHandAttr.Value == rightHandAttr.Value {
			return config.NO_SET_TWO_SAME_MATCHING_HANDS_SCALER, completedSetName, 2, mainMatchingTraits
		}

	} else if !hasCompletedSet {
		if matchingSetHandsCount == 2 && leftHandAttr.Value == rightHandAttr.Value {
			return config.INCOMPLETE_SET_TWO_SAME_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
	} else if hasCompletedSet {
		if matchingSetHandsCount == 2 && leftHandAttr.Value == rightHandAttr.Value {
			return config.INCOMPLETE_SET_TWO_SAME_MATCHING_HANDS_SCALER, completedSetName, matchingSetHandsCount, mainMatchingTraits
		}
	}
	return 1, "", 0, mainMatchingTraits
}

// calculateCompleteSets iterates over polymorph's attributes.
//
// Return if set has been completed, main set name, main set attrbiutes, secondary set name, secondary set attributes
func calculateCompleteSets(attributes []structs.Attribute) (bool, string, []string, string, []string) {
	var hasCompletedSet bool
	var mainSet int
	var mainSetName string

	setMap := make(map[string]int)
	setTraitsMap := make(map[string][]string)

	for _, attr := range attributes {
		for _, set := range attr.Sets {
			setMap[set]++
			setTraitsMap[set] = append(setTraitsMap[set], attr.TraitType)
			if setMap[set] == config.CombosMap[set] {
				hasCompletedSet = true
				mainSetName = set
				mainSet = setMap[set]
			}
		}
	}

	if mainSet == 0 {
		for k, v := range setMap {
			if v >= 2 && mainSet < v {
				mainSetName, mainSet = k, v
			}
		}
	}

	mainMatchingTraits := setTraitsMap[mainSetName]

	var secondarySetCount int
	var secondarySetName string

	for k, v := range setMap {
		if v >= 2 && secondarySetCount < v && k != mainSetName {
			possibleSecTraits := setTraitsMap[k]
			// Prevent same set from being added as a secondary set
			if !((helpers.StringInSlice("Left-Hand", mainMatchingTraits) || helpers.StringInSlice("Right-Hand", mainMatchingTraits)) &&
				len(possibleSecTraits) == 2 &&
				(possibleSecTraits[0] == "Right-Hand" || possibleSecTraits[0] == "Left-Hand") &&
				(possibleSecTraits[1] == "Right-Hand" || possibleSecTraits[1] == "Left-Hand")) {
				secondarySetName, secondarySetCount = k, v
			}
		}
	}
	secondaryMatchingTraits := setTraitsMap[secondarySetName]

	// It would be bad to have degen as main set while you have secondary set with the same number of traits
	// if len(mainMatchingTraits) == len(secondaryMatchingTraits) && mainSetName == "Party Degen" {
	// 	mainSetName, secondarySetName = secondarySetName, mainSetName
	// 	mainMatchingTraits, secondaryMatchingTraits = secondaryMatchingTraits, mainMatchingTraits
	// }

	return hasCompletedSet, mainSetName, mainMatchingTraits, secondarySetName, secondaryMatchingTraits
}
