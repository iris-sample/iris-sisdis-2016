package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	filetype "gopkg.in/h2non/filetype.v0"
	iris "gopkg.in/kataras/iris.v4"
)

type JSONResp struct {
	IsiBerkas    string `json:"isi_berkas"`
	LokasiBerkas string `json:"lokasi_berkas"`
	UkuranBerkas string `json:"ukuran_berkas"`
}

func hw4klienviewimage(c *iris.Context) {
	fileName := c.Param("name")
	r, _ := http.Get("https://prakash.sisdis.ui.ac.id/tugas4/server/getImage/" + fileName)
	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	jsonD := []byte(string(data))
	var q JSONResp
	json.Unmarshal(jsonD, &q)
	if q.IsiBerkas != "" {
		buf, _ := ioutil.ReadFile(q.LokasiBerkas)
		kind, _ := filetype.Match(buf)
		notError := true
		filePath := q.LokasiBerkas
		mime := kind.MIME.Value
		base64 := q.IsiBerkas
		fileSize := q.UkuranBerkas
		c.MustRender("hw4viewimage.html", struct {
			NotError bool
			FilePath string
			MIME     string
			Base64   string
			FileSize string
		}{NotError: notError, FilePath: filePath, MIME: mime, Base64: base64, FileSize: fileSize})
	} else {
		c.MustRender("hw4viewimage.html", nil)
	}
}
