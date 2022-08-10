package metadata

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"rarity-backend/constants"
	"rarity-backend/structs"
)

type Genome string
type Gene int

func (g Gene) toPath() string {
	if g < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(int(g)))
	}

	return strconv.Itoa(int(g))
}

func getGeneInt(g string, start, end, count int) int {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	return gene % count
}

func getEyeRightGene(g string) int {
	return getGeneInt(g, constants.EYE_RIGHT_START_IDX, constants.EYE_RIGHT_END_IDX, constants.EYE_RIGHT)
}
func GetEyeRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getEyeRightGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.EyeRight,
		Value:     trait,
	}
}
func getEyeRightPath(g string) string {
	gene := getEyeRightGene(g)
	return Gene(gene).toPath()
}

func getEyeLeftGene(g string) int {
	return getGeneInt(g, constants.EYE_LEFT_START_IDX, constants.EYE_LEFT_END_IDX, constants.EYE_LEFT)
}
func GetEyeLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getEyeLeftGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.EyeLeft,
		Value:     trait,
	}
}
func getEyeLeftPath(g string) string {
	gene := getEyeLeftGene(g)
	return Gene(gene).toPath()
}

func getBeardBottomRightGene(g string) int {
	return getGeneInt(g, constants.BEARD_BOTTOM_RIGHT_START_IDX, constants.BEARD_BOTTOM_RIGHT_END_IDX, constants.BEARD_BOTTOM_RIGHT)
}
func GetBeardBottomRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBeardBottomRightGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.BeardBottomRight,
		Value:     trait,
	}
}
func getBeardBottomRightPath(g string) string {
	gene := getBeardBottomRightGene(g)
	return Gene(gene).toPath()
}

func getBeardBottomLeftGene(g string) int {
	return getGeneInt(g, constants.BEARD_BOTTOM_LEFT_START_IDX, constants.BEARD_BOTTOM_LEFT_END_IDX, constants.BEARD_BOTTOM_LEFT)
}
func GetBeardBottomLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBeardBottomLeftGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.BeardBottomLeft,
		Value:     trait,
	}
}
func getBeardBottomLeftPath(g string) string {
	gene := getBeardBottomLeftGene(g)
	return Gene(gene).toPath()
}

func getLipsRightGene(g string) int {
	return getGeneInt(g, constants.LIPS_RIGHT_START_IDX, constants.LIPS_RIGHT_END_IDX, constants.LIPS_RIGHT)
}

func GetLipsRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getLipsRightGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.LipsRight,
		Value:     trait,
	}
}
func getLipsRightPath(g string) string {
	gene := getLipsRightGene(g)
	return Gene(gene).toPath()
}

func getLipsLeftGene(g string) int {
	return getGeneInt(g, constants.LIPS_LEFT_START_IDX, constants.LIPS_LEFT_END_IDX, constants.LIPS_LEFT)
}

func GetLipsLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getLipsLeftGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.LipsLeft,
		Value:     trait,
	}
}
func getLipsLeftPath(g string) string {
	gene := getLipsLeftGene(g)
	return Gene(gene).toPath()
}

func getBeardTopRightGene(g string) int {
	return getGeneInt(g, constants.BEARD_TOP_RIGHT_START_IDX, constants.BEARD_TOP_RIGHT_END_IDX, constants.BEARD_TOP_RIGHT)
}

func GetBeardTopRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBeardTopRightGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.BeardTopRight,
		Value:     trait,
	}
}
func getBeardTopRightPath(g string) string {
	gene := getBeardTopRightGene(g)
	return Gene(gene).toPath()
}

func getBeardTopLeftGene(g string) int {
	return getGeneInt(g, constants.BEARD_TOP_LEFT_START_IDX, constants.BEARD_TOP_LEFT_END_IDX, constants.BEARD_TOP_LEFT)
}
func GetBeardTopLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBeardTopLeftGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.BeardTopLeft,
		Value:     trait,
	}
}
func getBeardTopLeftPath(g string) string {
	gene := getBeardTopLeftGene(g)
	return Gene(gene).toPath()
}

func getEarsRightGene(g string) int {
	return getGeneInt(g, constants.EAR_RIGHT_START_IDX, constants.EAR_RIGHT_END_IDX, constants.EAR_RIGHT)
}

func GetEarsRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getEarsRightGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.EarRight,
		Value:     trait,
	}
}

func getEarsRightPath(g string) string {
	gene := getEarsRightGene(g)
	return Gene(gene).toPath()
}

func getEarsLeftGene(g string) int {
	return getGeneInt(g, constants.EAR_LEFT_START_IDX, constants.EAR_LEFT_END_IDX, constants.EAR_LEFT)
}

func GetEarsLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getEarsLeftGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.EarLeft,
		Value:     trait,
	}
}

func getEarsLeftPath(g string) string {
	gene := getEarsLeftGene(g)
	return Gene(gene).toPath()
}

func getHairRightGene(g string) int {
	return getGeneInt(g, constants.HAIR_RIGHT_START_IDX, constants.HAIR_RIGHT_END_IDX, constants.HAIR_RIGHT)
}

func GetHairRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getHairRightGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.HairRight,
		Value:     trait,
	}
}

func getHairRightPath(g string) string {
	gene := getHairRightGene(g)
	return Gene(gene).toPath()
}

func getHairLeftGene(g string) int {
	return getGeneInt(g, constants.HAIR_LEFT_START_IDX, constants.HAIR_LEFT_END_IDX, constants.HAIR_LEFT)
}

func GetHairLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getHairLeftGene(g)
	trait := configService.Traits[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.HairLeft,
		Value:     trait,
	}
}

func getHairLeftPath(g string) string {
	gene := getHairLeftGene(g)
	return Gene(gene).toPath()
}

func getBackgroundGene(g string) int {
	return getGeneInt(g, constants.BACKGROUND_GENE_START_IDX, constants.BACKGROUND_GENE_END_IDX, constants.BACKGROUND_GENE_COUNT)
}

func GetBackgroundGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBackgroundGene(g)
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Background,
		Value:     configService.Traits[gene],
	}
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(tokenId string) string {
	return fmt.Sprintf("Polymorphic Face #%v", tokenId)
}

func (g *Genome) description(tokenId string) string {
	return fmt.Sprintf("The Polymorphic Face #%v has a unique genetic code! You can scramble your Polymorphic face at anytime.", tokenId)
}

func (g *Genome) Genes() []string {
	gStr := string(*g)

	res := make([]string, 0, constants.GENES_COUNT)

	res = append(res, getBeardBottomRightPath(gStr))
	res = append(res, getBeardBottomLeftPath(gStr))
	res = append(res, getLipsRightPath(gStr))
	res = append(res, getLipsLeftPath(gStr))
	res = append(res, getBeardTopRightPath(gStr))
	res = append(res, getBeardTopLeftPath(gStr))
	res = append(res, getEyeRightPath(gStr))
	res = append(res, getEyeLeftPath(gStr))
	res = append(res, getEarsRightPath(gStr))
	res = append(res, getEarsLeftPath(gStr))
	res = append(res, getHairRightPath(gStr))
	res = append(res, getHairLeftPath(gStr))
	res = append(res, getBackgroundGenePath(gStr))

	return res
}

func (g *Genome) attributes(configService *structs.ConfigService) []structs.Attribute {
	gStr := string(*g)

	res := make([]structs.Attribute, 0, constants.GENES_COUNT)
	res = append(res, GetBeardBottomRightGeneAttribute(gStr, configService))
	res = append(res, GetBeardBottomLeftGeneAttribute(gStr, configService))
	res = append(res, GetLipsRightGeneAttribute(gStr, configService))
	res = append(res, GetLipsLeftGeneAttribute(gStr, configService))
	res = append(res, GetBeardTopRightGeneAttribute(gStr, configService))
	res = append(res, GetBeardTopLeftGeneAttribute(gStr, configService))
	res = append(res, GetEyeRightGeneAttribute(gStr, configService))
	res = append(res, GetEyeLeftGeneAttribute(gStr, configService))
	res = append(res, GetEarsRightGeneAttribute(gStr, configService))
	res = append(res, GetEarsLeftGeneAttribute(gStr, configService))
	res = append(res, GetHairRightGeneAttribute(gStr, configService))
	res = append(res, GetHairLeftGeneAttribute(gStr, configService))
	res = append(res, GetBackgroundGeneAttribute(gStr, configService))
	return res
}

func (g *Genome) Metadata(tokenId string, configService *structs.ConfigService) structs.Metadata {
	var m structs.Metadata
	m.Attributes = g.attributes(configService)
	m.Name = g.name(tokenId)
	m.Description = g.description(tokenId)
	m.ExternalUrl = fmt.Sprintf("%s%s", constants.EXTERNAL_URL, tokenId)

	genes := g.Genes()
	facesImgUrl := os.Getenv("FACES_IMAGE_URL")
	imageUrl := strings.Builder{}
	imageUrl.WriteString(facesImgUrl)

	// imageUrl3D := strings.Builder{}
	// imageUrl3D.WriteString(constants.POLYMORPH_IMAGE_URL_3D)

	for _, gene := range genes {
		imageUrl.WriteString(gene)
		// imageUrl3D.WriteString(gene)
	}

	imageUrl.WriteString(".jpg")
	// imageUrl3D.WriteString(".jpg")

	m.Image = imageUrl.String()
	// m.Image3D = imageUrl3D.String()

	return m
}
