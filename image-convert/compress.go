package imageconvert

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"os/exec"
)

func CompressImg(imageWithExt string, imgUrl string) (resWebp *Webp) {
	var webp Webp
	originalImagePath := "/home/original-image/" + imageWithExt
	imageBytes, err := os.ReadFile(originalImagePath)

	if err != nil {
		fmt.Println(err)
	}
	contentType := http.DetectContentType(imageBytes)
	fmt.Println(contentType)

	switch contentType {
	case "image/png":
		// cmd = exec.Command("oxipng", "-o", "4", "-i", "1", "--strip", "safe", "/var/www/html/snapbox.me/public/pngOut/"+outFilePath)
		// fmt.Println(cmd)
		cmd := exec.Command("pngquant", "--quality=70-99", "/home/original-image/"+imageWithExt, "--output", "/var/www/html/snapbox.me/public/compress-out/"+imageWithExt)
		errC := cmd.Run()
		if errC != nil {
			webp.WebpConverted = 0
		}
		webp.WebpConverted = 1
		webp.ImagePath = imageWithExt
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			fmt.Println(err)
		}

		out, err := os.Create("/var/www/html/snapbox.me/public/compress-out/" + imageWithExt)
		defer out.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var opt jpeg.Options
		opt.Quality = 85
		if err := jpeg.Encode(out, img, &opt); err != nil {
			fmt.Println(err)
		}
		webp.WebpConverted = 1
		webp.ImagePath = imageWithExt
	}
	fmt.Println(webp)
	return &webp
}
