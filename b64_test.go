package base64img

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestJpg(t *testing.T) {
	t.Run("name without extension", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img.txt")
		jpg := new(Jpg)
		jpg.Filename = "test"
		jpg.DataUrl = string(img)
		err := jpg.CheckExtension().Build()
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		t.Log("success")
	})
	t.Run("name with extension jpg", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img.txt")
		jpg := new(Jpg)
		jpg.Filename = "buku.jpg"
		jpg.DataUrl = string(img)
		err := jpg.CheckExtension().Build()
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		t.Log("success")
	})
	t.Run("name with extension jpeg", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img.txt")
		jpg := new(Jpg)
		jpg.Filename = "buku.jpeg"
		jpg.DataUrl = string(img)
		err := jpg.CheckExtension().Build()
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		t.Log("success")
	})
	t.Run("file name is invalid", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img.txt")
		jpg := new(Jpg)
		jpg.Filename = ".jpeg"
		jpg.DataUrl = string(img)
		err := jpg.CheckExtension().Build()
		fmt.Println(err)
		if err == nil {
			t.FailNow()
		}
		t.Log("success")
	})
}

func TestPng(t *testing.T) {
	t.Run("name without extension", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img_png.txt")
		png := new(Png)
		png.Filename = "test"
		png.DataUrl = string(img)
		err := png.CheckExtension().Build()
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		t.Log("success")
	})
	t.Run("name with extension png", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img_png.txt")
		png := new(Png)
		png.Filename = "buku.png"
		png.DataUrl = string(img)
		err := png.CheckExtension().Build()
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		t.Log("success")
	})
	t.Run("invalid file name", func(t *testing.T) {
		img,_ := ioutil.ReadFile("img_png.txt")
		png := new(Png)
		png.Filename = ".png"
		png.DataUrl = string(img)
		err := png.CheckExtension().Build()
		if err == nil {
			fmt.Println(err.Error())
			t.FailNow()
		}
		t.Log("success")
	})
}
