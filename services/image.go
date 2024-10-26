package services

import (
	"fmt"
	"image/png"
	"os"
	"prelo/database"
	"prelo/models"

	"github.com/nfnt/resize"
)

// MakeImages -
func MakeImages(image *models.Image, id string) {

	thumbPath := ResizeImage(image.Path, "thumb", 300)
	thumbURL := UploadToSpaces("items/thumbs", thumbPath, id) // := UploadToSpaces("properties/thumbs", thumbPath, image.ID.String())
	fmt.Println(thumbPath)

	mainPath := ResizeImage(image.Path, "", 1000)
	imageURL := UploadToSpaces("items", mainPath, image.ID.String())
	db := database.DB.Db
	db.Model(&image).Updates(models.Image{
		URL:      imageURL,
		ThumbURL: thumbURL,
	})
	if image.ImageType == "user" {
		//db.Exec("update users set picture =?  , avatar=? where id=?", thumbURL, thumbURL, id)
	}
	if image.ImageType == "item" {
		//db.Exec("update properties set picture =?  where id=?", image.URL, id)
	}

	//image.Url = services.UploadToSpaces("properties", fmt.Sprintf("%s-%s", image.ID.String(), file.Filename))
	os.Remove(image.Path)
}

// ResizeImage -
func ResizeImage(path, imageType string, width uint) string {
	// open "test.jpg"
	fmt.Println("resise: ", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("oen", err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)

	if err != nil {
		fmt.Println("decode", err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, 0, img, resize.Lanczos3)

	resizedPath := path + "-" + imageType + ".png"

	out, err := os.Create(resizedPath)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)

	//fmt.Println("resized completed!")
	return resizedPath
}
