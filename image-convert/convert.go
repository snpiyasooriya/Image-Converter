package imageconvert

import (
	"os/exec"
	"path"
	"strings"
)

func ConvertToGif(imagePath string, imgUrl string) (resWebp *Webp) {
	var webp Webp
	outFilePath := strings.Replace(imagePath, path.Ext(imgUrl), ".gif", 1)
	cmd := exec.Command("/usr/bin/python3", "vdo2gif.py", "/home/original-video/"+imagePath, "/var/www/html/snapbox.me/public/gif-out/"+outFilePath)
	errWebp := cmd.Run()
	if errWebp != nil {
		webp.WebpConverted = 0
	}
	webp.WebpConverted = 1
	webp.ImagePath = outFilePath
	return &webp
}
