package helpers

import (
	"basic-trade-api/configs"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(FileHeader *multipart.FileHeader, fileName string) (string, error) {
	// Initialize context for Cloudinary
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// After function is finished, cancel the context
	defer cancel()

	// Add cloudinary instance from provided parameters
	cld, err := configs.InitializeCloudinary()
	if err != nil {
		log.Fatalln("Failed to add cloudinary instance")
		return "", err
	}

	// Convert fileHeader into fileReader
	fileReader, err := convertFile(FileHeader)
	if err != nil {
		return "", err
	}

	// Retrieve cloudinary folder
	cldFolder := os.Getenv("CLOUDINARY_FOLDER")
	if cldFolder == "" {
		log.Fatalln("Cloudinary folder name is empty. Check the .env file.")
		return "", err
	}
	// Prepare uploadParams
	uploadParams := uploader.UploadParams{
		PublicID: fileName,
		Folder:   cldFolder,
	}

	// Upload file
	uploadResult, err := cld.Upload.Upload(c, fileReader, uploadParams)
	if err != nil {
		log.Fatalln("Failed uploading file to cloudinary:", err)
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func DeleteFile(imageURL string) error {
	// Initialize context for Cloudinary
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// After function is finished, cancel the context
	defer cancel()

	// Add cloudinary instance from provided parameters
	cld, err := configs.InitializeCloudinary()
	if err != nil {
		log.Fatalln("Failed to add cloudinary instance")
		return err
	}

	// Delete asset
	deletedResult, err := cld.Admin.DeleteAssets(c, admin.DeleteAssetsParams{
		PublicIDs:    []string{retrievePublicID(imageURL)},
		DeliveryType: "upload",
		AssetType:    "image",
	})
	if err != nil {
		log.Default().Fatalln("Failed to delete asset in cloudinary")
		return err
	}

	fmt.Println("Deleted", deletedResult.DeletedCounts, "asset from cloudinary")

	return nil
}

// Convert from FileHeader into bytes.Reader
func convertFile(fileHeader *multipart.FileHeader) (*bytes.Reader, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		log.Fatalln("Failed to open the file header:", err)
		return nil, err
	}
	// Close the file after the function is finished
	defer file.Close()

	// Create buffer to read the file into an in-memory buffer
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, file)
	if err != nil {
		log.Fatalln("Failed copying file into in-memory buffer:", err)
		return nil, err
	}

	// Create a bytes.Reader from the buffer
	fileReader := bytes.NewReader(buffer.Bytes())
	return fileReader, nil
}

// Remove file extension, extract only the file name
func RemoveExtension(filename string) string {
	return path.Base(filename[:len(filename)-len(path.Ext(filename))])
}

// Retrieve public ID from URL
func retrievePublicID(imageURL string) string {
	fileName := RemoveExtension(path.Base(imageURL))
	return "basic-trade-api/" + fileName
}
