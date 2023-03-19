package imageconvert

import (
	"fmt"
	"math/rand"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Webp struct {
	WebpConverted int    `json:"webpConvert"`
	ImagePath     string `json:"imagePath"`
}

func (webp *Webp) WebpConvert(imgUrl string, cmdType string) {
	fileName := strconv.Itoa(rand.Intn(1000000))
	fileName += path.Base(imgUrl)
	outFilePath := strings.Replace(fileName, path.Ext(imgUrl), ".webp", 1)

	_, err := GrabImg(imgUrl, fileName, cmdType)
	if err != nil {
		fmt.Println(err)
	}
	if cmdType == "convert-to-webp" {
		cmd := exec.Command("cwebp", "-q", "50", "/home/png/"+fileName, "-o", "/var/www/html/snapbox.me/public/webp-out/"+outFilePath)
		errWebp := cmd.Run()
		if errWebp != nil {
			webp.WebpConverted = 0
		}
		webp.WebpConverted = 1
		webp.ImagePath = outFilePath
	}
	if cmdType == "compress-img" {
		resWebp := CompressImg(fileName, imgUrl)
		webp.WebpConverted = resWebp.WebpConverted
		webp.ImagePath = resWebp.ImagePath
	}
	if cmdType == "convert-to-gif" {
		resWebp := ConvertToGif(fileName, imgUrl)
		webp.WebpConverted = resWebp.WebpConverted
		webp.ImagePath = resWebp.ImagePath
	}
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
