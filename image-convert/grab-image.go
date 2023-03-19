package imageconvert

import (
	"fmt"

	"github.com/imroc/req/v3"
)

func GrabImg(url string, fileName string, cmdType string) (imgFile string, resErr error) {

	client := req.C()
	var err error
	var t *req.Response
	if cmdType == "convert-to-webp" {
		t, err = client.R().SetOutputFile("/home/png/" + fileName).Get(url)

	}
	if cmdType == "compress-png" {
		t, err = client.R().SetOutputFile("/home/png/" + fileName).Get(url)

	}
	if cmdType == "jpeg-to-png" {
		t, err = client.R().SetOutputFile("/home/png/" + fileName).Get(url)
		fmt.Println(fileName)

	}
	if cmdType == "compress-img" {
		t, err = client.R().SetOutputFile("/home/original-image/" + fileName).Get(url)
		fmt.Println(fileName)

	}
	if cmdType == "convert-to-gif" {
		t, err = client.R().SetOutputFile("/home/original-video/" + fileName).Get(url)
		fmt.Println(fileName)

	}
	if err != nil {
		return "", err
	}

	return t.String(), nil

}
