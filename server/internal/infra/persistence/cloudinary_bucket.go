package persistence

import (
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary(cloudName, apiKey, apiSecret string) *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		fmt.Println("Failed to initialize Cloudinary:", err)
		return nil
	}

	return cld
}
