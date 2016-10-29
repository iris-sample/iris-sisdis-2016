package main

import (
	"bufio"
	"encoding/base64"
	"os"
	"path/filepath"

	humanize "github.com/dustin/go-humanize"
	"github.com/kataras/iris"
)

func hw4servergetimage(c *iris.Context) {
	fileName := c.Param("name")
	if fileName != "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		imgFile, err := os.Open("uploads/" + fileName)
		defer imgFile.Close()
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"error": "File Not Found"})
		} else {
			var fileType string
			fileType = getFormat(imgFile)
			if fileType != "" {
				var size int64
				fInfo, _ := imgFile.Stat()
				size = fInfo.Size()
				buf := make([]byte, size)
				fReader := bufio.NewReader(imgFile)
				fReader.Read(buf)
				imgBase64Str := base64.StdEncoding.EncodeToString(buf)
				var resp JSONResp
				resp.IsiBerkas = imgBase64Str
				resp.LokasiBerkas = dir + "/" + imgFile.Name()
				resp.UkuranBerkas = humanize.Bytes(uint64(size))
				c.JSON(iris.StatusOK, resp)
			} else {
				c.JSON(iris.StatusOK, iris.Map{"error": "Not an Image"})
			}
		}
	} else {
		c.JSON(iris.StatusOK, iris.Map{"error": "Not Allowed"})
	}
}
