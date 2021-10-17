package base64img

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/png"
	"os"
	"strings"
)

type Png struct {
	Filename  string
	DataUrl   string
	extension string
	base64img string
}

func (p *Png) CheckExtension() *Png {
	var valid bool = strings.Contains(p.DataUrl, "data:image/png")
	if valid == false {
		panic(errors.New("image is not png"))
	}
	items := strings.Split(p.DataUrl, ";")
	p.extension = strings.Split(items[0], "/")[1]
	p.DataUrl = strings.Split(items[1], ",")[1]
	return p
}

func (p Png) Build() error {
	var fileName string
	decoded, err := base64.StdEncoding.DecodeString(p.DataUrl)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(decoded)
	pngImg,err := png.Decode(reader)
	if err != nil {
		return err
	}
	if strings.Contains(p.Filename,".png") {
		items := strings.Split(p.Filename,".")
		if len(items[0]) > 0 {
			fileName = p.Filename
		} else {
			return errors.New("file name is invalid")
		}
	} else {
		fileName = fmt.Sprintf("%s.%s", p.Filename, p.extension)
	}
	file,err := os.Create(fileName)
	if err != nil {
		return err
	}
	err = png.Encode(file,pngImg)
	return err
}

