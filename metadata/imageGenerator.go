package metadata

import (
	"context"
	"fmt"
	"image"
	"image/color"
	"log"
	"net/http"
	"rarity-backend/constants"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/disintegration/imaging"
)

const IMG_SIZE = 4000

func CheckImageAndCreate(imageURL string, genes []string) {
	imageExists := ImageExists(imageURL)
	if !imageExists {
		generateAndSaveImage(genes)
	}
}

func ImageExists(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
		return false
	}

	return resp.StatusCode != 404
}

func combineRemoteImages(bucket *storage.BucketHandle, basePath string, overlayPaths ...string) (*image.NRGBA, error) {

	ctx := context.Background()

	baseReader, err := bucket.Object(basePath).NewReader(ctx)
	if err != nil {
		// log.Fatalf("failed to open image: %v", err)
		log.Printf("failed to open image: %v", err)
		return nil, err
	}
	defer baseReader.Close()
	base, err := imaging.Decode(baseReader)
	if err != nil {
		// log.Fatalf("failed to decode image: %v", err)
		log.Printf("failed to decode image: %v", err)
		return nil, err
	}
	dst := imaging.New(IMG_SIZE, IMG_SIZE, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, base, image.Pt(0, 0))

	for _, op := range overlayPaths {
		r, err := bucket.Object(op).NewReader(ctx)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		defer r.Close()
		o, err := imaging.Decode(r)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		dst = imaging.Overlay(dst, o, image.Pt(0, 0), 1)
	}

	return dst, nil
}

func ReverseGenesOrder(genes []string) []string {
	res := make([]string, 0, len(genes))
	for i := len(genes) - 1; i >= 0; i-- {
		res = append(res, genes[i])
	}
	return res
}

func saveToGCloud(i *image.NRGBA, name string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Println("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(constants.GCLOUD_UPLOAD_BUCKET_NAME).Object(name).NewWriter(ctx)
	// f, err := imaging.FormatFromFilename(name)
	// if err != nil {
	// 	log.Errorf("Format from filename: %v", err)
	// }
	err = imaging.Encode(bucket, i, imaging.JPEG, imaging.JPEGQuality(80))

	if err != nil {
		log.Println("Upload: %v", err)
	}

	if err = bucket.Close(); err != nil {
		log.Println("Writer.Close: %v", err)
	}

}

func generateAndSaveImage(genes []string) {
	// Reverse
	revGenes := ReverseGenesOrder(genes)

	f := make([]string, len(genes))

	for i, gene := range revGenes {
		f[i] = fmt.Sprintf("./images/%v/%s.png", i, gene)
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Println("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(constants.GCLOUD_SOURCE_BUCKET_NAME)

	i, err := combineRemoteImages(bucket, f[0], f[1:]...)
	if err != nil {
		b := strings.Builder{}

		for _, gene := range genes {
			b.WriteString(gene)
		}

		b.WriteString(".jpg") // Finish with jpg extension

		saveToGCloud(i, b.String())

	}
}
