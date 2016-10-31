package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	iris "gopkg.in/kataras/iris.v4"
)

type JSONReq struct {
	NamaBerkas string `json:"nama_berkas"`
	IsiBerkas  string `json:"isi_berkas"`
}

type JSONResult struct {
	Result string `json:"result"`
}

func hw4getuploadimage(c *iris.Context) {
	c.MustRender("hw4uploadimage.html", nil)
}

func hw4postuploadimage(c *iris.Context) {
	var result string
	req, _ := c.FormFile("file")
	fileName := req.Filename
	src, err := req.Open()
	buff, err := ioutil.ReadAll(src)
	base64Str := base64.StdEncoding.EncodeToString(buff)
	var request JSONReq
	request.IsiBerkas = base64Str
	request.NamaBerkas = fileName
	out, err := json.Marshal(request)
	r, err := http.Post("https://prakash.sisdis.ui.ac.id/tugas4/server/postImage", "application/json", bytes.NewBuffer(out))
	if err != nil {
		result = "Not connected."
	} else {
		data, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		status := []byte(string(data))
		var q JSONResult
		json.Unmarshal(status, &q)
		result = "Result : " + q.Result
	}
	c.MustRender("hw4uploadimage.html", struct {
		Result string
	}{Result: result})
}
