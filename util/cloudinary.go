package util

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-01/group-project/pejuang-rupiah/backend/config"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := config.GetEnv("CLOUDINARY_SECRET")
	cldName := config.GetEnv("CLOUDINARY_NAME")
	cldKey := config.GetEnv("CLOUDINARY_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}

func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
	ctx := context.Background()
	cld, err := SetupCloudinary()
	if err != nil {
		return "", err
	}
	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}
	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}

func GetVariantImageURL(imageId string) (string, error) {
	file, err := os.Open(fmt.Sprintf("./imageuploads/%v.jpeg", imageId))
	if err != nil {
		return "", err
	}

	defer file.Close()

	currentTime := time.Now().UnixNano()
	fileExtension := path.Ext(file.Name())
	originalFilename := file.Name()[:len(file.Name())-len(fileExtension)]
	newFilename := fmt.Sprintf("%s_%d", originalFilename, currentTime)
	fileName := strings.Split(newFilename, "./imageuploads/")

	imageURL, err := UploadToCloudinary(file, fileName[1])
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func ConvertToJPEG(file multipart.File) (io.Reader, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	var convertedImageBuffer bytes.Buffer

	options := &jpeg.Options{Quality: 50}
	if err := jpeg.Encode(&convertedImageBuffer, img, options); err != nil {
		return nil, err
	}

	convertedImageReader := bytes.NewReader(convertedImageBuffer.Bytes())

	return convertedImageReader, nil
}
