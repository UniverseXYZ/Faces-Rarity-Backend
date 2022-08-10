package helpers

// DetectGeneDifferences compares two gene string and identifies the number of differences and the index of the differences.
//
// The difference index is later used in GetAttribute to calculate which attribute has changed
func DetectGeneDifferences(oldGene string, newGene string) (int, int) {
	oldGene = ReverseString(oldGene)
	newGene = ReverseString(newGene)
	var differences, geneIndex int
	if oldGene == "0" {
		return 0, 0
	}
	// bigInt.String() removes leading zeroes so we have to recover them
	if len(oldGene) != len(newGene) {
		if len(oldGene) < len(newGene) {
			lenDiff := len(newGene) - len(oldGene)
			for i := 0; i < lenDiff; i++ {
				oldGene = "0" + oldGene
			}
		} else {
			lenDiff := len(oldGene) - len(newGene)
			for i := 0; i < lenDiff; i++ {
				newGene = "0" + newGene
			}

		}
	}

	for i := range oldGene {
		if oldGene[i] != newGene[i] {
			differences++
			geneIndex = i
		}
	}

	return geneIndex, differences
}
