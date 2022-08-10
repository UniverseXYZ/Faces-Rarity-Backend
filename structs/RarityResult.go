package structs

type RarityResult struct {
	HasCompletedSet       bool
	MainSetName           string
	MainMatchingTraits    []string
	SecSetName            string
	SecMatchingTraits     []string
	ColorMismatches       int
	HandsSetName          string
	HandsScaler           float64
	MatchingHands         int
	NoColorMismatchScaler float64
	ColorMismatchScaler   float64
	VirginScaler          float64
	BaseRarity            float64
	ScaledRarity          float64
}
