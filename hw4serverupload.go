package main

import (
	"bytes"
	"encoding/base64"
	"os"
	"strings"

	"github.com/kataras/iris"
)

func hw4serveruploadimage(c *iris.Context) {
	req := new(JSONReq)
	res := new(JSONResult)
	c.ReadJSON(req)
	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(req.IsiBerkas))
	f, err := os.Open("uploads/" + req.NamaBerkas)
	defer f.Close()
	if err != nil {
		imgSave, _ := os.Create("uploads/" + req.NamaBerkas)
		defer imgSave.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(imageReader)
		file, _ := imgSave.Write(buf.Bytes())
		_ = file
		fileType := getFormat(imgSave)
		if fileType != "" {
			res.Result = "Success"
		} else {
			_ = os.Remove("uploads/" + req.NamaBerkas)
			res.Result = "File Not An Image"
		}
	} else {
		res.Result = "File Exist"
	}
	c.JSON(iris.StatusOK, res)
}

func getFormat(file *os.File) string {
	bytes := make([]byte, 4)
	n, _ := file.ReadAt(bytes, 0)
	if n < 4 {
		return ""
	}
	if bytes[0] == 0x89 && bytes[1] == 0x50 && bytes[2] == 0x4E && bytes[3] == 0x47 {
		return "png"
	}
	if bytes[0] == 0xFF && bytes[1] == 0xD8 {
		return "jpg"
	}
	if bytes[0] == 0x47 && bytes[1] == 0x49 && bytes[2] == 0x46 && bytes[3] == 0x38 {
		return "gif"
	}
	if bytes[0] == 0x42 && bytes[1] == 0x4D {
		return "bmp"
	}
	return ""
}
