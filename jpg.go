package base64img

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"os"
	"strings"
)

type Jpg struct {
	Filename  string
	DataUrl   string
	extension string
	base64img string
}

func (j *Jpg) CheckExtension() *Jpg {
	var valid bool = strings.Contains(j.DataUrl, "data:image/jpeg")
	fmt.Println(valid)
	if valid == false {
		valid = strings.Contains(j.DataUrl, "data:image/jpg")
		if valid == false {
			panic(errors.New("extension is not jpeg or jpg"))
		}
	}
	items := strings.Split(j.DataUrl, ";")
	extension := strings.Split(items[0], "/")[1]
	encodedImg := strings.Split(items[1], ",")[1]
	j.extension = extension
	j.base64img = encodedImg
	return j
}

func (j *Jpg) Build() error {
	var fileName string
	decodedImg, err := base64.StdEncoding.DecodeString(j.base64img)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(decodedImg)
	jpg, err := jpeg.Decode(reader)
	if err != nil {
		return err
	}
	if strings.Contains(j.Filename, ".jpg") || strings.Contains(j.Filename,".jpeg") {
		items := strings.Split(j.Filename,".")
		if len(items[0]) > 0 {
			fileName = j.Filename
		} else {
			return errors.New("file name is invalid")
		}
	} else {
		fileName = fmt.Sprintf("%s.%s", j.Filename, j.extension)
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	err = jpeg.Encode(file, jpg, &jpeg.Options{Quality: 75})
	return err
}

