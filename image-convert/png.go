package imageconvert

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

// ToPng converts an image to png
func ToPng(imagePath string, outFilePath string) {

	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println(err)
	}
	contentType := http.DetectContentType(imageBytes)
	fmt.Println(contentType)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			fmt.Println(err)
		}

		out, err := os.Create("/var/www/html/snapbox.me/public/png-out/" + outFilePath)
		defer out.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := png.Encode(out, img); err != nil {
			fmt.Println(err)
		}
	}

}
