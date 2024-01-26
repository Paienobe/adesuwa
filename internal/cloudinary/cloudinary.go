package cloudinary

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(file multipart.File, publicId string) string {
	// Add your Cloudinary product environment credentials.
	cld, _ := cloudinary.New()

	// Upload the my_picture.jpg image and set the PublicID to "my_image".
	var ctx = context.Background()
	uploadParams := uploader.UploadParams{PublicID: publicId}
	resp, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		fmt.Println("Error uploading image:", err)
		return ""
	}
	return resp.SecureURL
	// log.Println("Image uploaded successfully. Secure URL:", resp.SecureURL)

	// // Re-upload the my_picture.jpg image with the same PublicID "my_image".
	// resp, err = cld.Upload.Upload(ctx, file, uploadParams)
	// if err != nil {
	// 	fmt.Println("Error re-uploading image:", err)
	// 	return
	// }
	// log.Println("Image re-uploaded successfully. Secure URL:", resp.SecureURL)

	// Get details about the image with PublicID "my_image" and log the secure URL.
	// assetResp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: publicId})
	// if err != nil {
	// 	fmt.Println("Error fetching asset details:", err)
	// 	return
	// }
	// log.Println("Asset details fetched successfully. Secure URL:", assetResp.SecureURL)
}
