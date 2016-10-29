package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kataras/iris"
)

type WSDLHw3 struct {
	Addr WSDLHw3Address `xml:"service>port>address"`
}

type WSDLHw3Address struct {
	URL string `xml:"location,attr"`
}

type RequestHello struct {
	XMLName      xml.Name `xml:"Envelope"`
	Header       string   `xml:"Header"`
	HelloRequest string   `xml:"Body>HelloRequest"`
}

func hw3clientget(c *iris.Context) {
	c.MustRender("hw3.html", nil)
}

func hw3clientpost(c *iris.Context) {
	var resp string
	url := c.FormValueString("url")
	kata := c.FormValueString("string")
	fmt.Print(url)
	fmt.Print(kata)
	if url == "" && kata == "" {
		resp = "Invalid WSDL URL and no string"
	} else if url == "" {
		resp = "Invalid WSDL URL"
	} else if kata == "" {
		resp = "No string"
	} else {
		r1, err := http.Get(url)
		if err != nil {
			resp = "Invalid WSDL URL"
		} else {
			data1, _ := ioutil.ReadAll(r1.Body)
			r1.Body.Close()
			xml1 := []byte(string(data1))
			var q1 WSDLHw3
			xml.Unmarshal(xml1, &q1)
			url := q1.Addr.URL
			buffer := new(RequestHello)
			buffer.HelloRequest = kata
			out, err := xml.Marshal(buffer)
			r2, err := http.Post(url, "text/xml", bytes.NewBuffer(out))
			if err != nil {
				resp = "Unexpected Error"
			} else {
				data2, _ := ioutil.ReadAll(r2.Body)
				r2.Body.Close()
				var q2 ReplyHello
				xml2 := []byte(string(data2))
				xml.Unmarshal(xml2, &q2)
				resp = q2.HelloResponse
			}
		}
	}
	c.MustRender("hw3.html", struct {
		Resp string
	}{Resp: resp})
}
