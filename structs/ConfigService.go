package structs

type AttributeSet struct {
	Name string   `json:"name"`
	Sets []string `json:"sets"`
}

type ConfigService struct {
	Background       []string       `json:"background"`
	EyeRight         []AttributeSet `json:"eyeright"`
	EyeLeft          []AttributeSet `json:"eyeleft"`
	BeardTopRight    []AttributeSet `json:"beardtopright"`
	BeardTopLeft     []AttributeSet `json:"beardtopleft"`
	BeardBottomRight []AttributeSet `json:"beardbottomright"`
	BeardBottomLeft  []AttributeSet `json:"beardbottomleft"`
	LipsRight        []AttributeSet `json:"lipsright"`
	LipsLeft         []AttributeSet `json:"lipsleft"`
	EarsRight        []AttributeSet `json:"earsright"`
	EarsLeft         []AttributeSet `json:"earsleft"`
	HairRight        []AttributeSet `json:"hairright"`
	HairLeft         []AttributeSet `json:"hairleft"`
	Traits           []string       `json:"traits"`
}
