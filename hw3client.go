package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	iris "gopkg.in/kataras/iris.v4"
)

type WSDLHw3 struct {
	Addr WSDLHw3Address `xml:"service>port>address"`
}

type WSDLHw3Address struct {
	URL string `xml:"location,attr"`
}

type RequestHello struct {
	HelloRequest string `xml:"Body>HelloRequest"`
}

func hw3clientget(c *iris.Context) {
	c.MustRender("hw3.html", nil)
}

func hw3clientpost(c *iris.Context) {
	var resp string
	url := c.FormValueString("url")
	kata := c.FormValueString("string")
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
			var data = `<soapenv:Envelope
				xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
				xmlns:hy="http://www.herongyang.com/Service/">
			<soapenv:Header/>
			<soapenv:Body>
				<hy:HelloRequest>`
			data += kata
			data += `</hy:HelloRequest>
				</soapenv:Body>
			</soapenv:Envelope>`
			r2, err := http.Post(url, "text/xml", bytes.NewBuffer([]byte(data)))
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
