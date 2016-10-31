package main

import (
	"encoding/xml"

	"github.com/kataras/iris"
)

type RPC struct {
	Body RequestBody `xml:"Body"`
}

type RequestBody struct {
	XMLName xml.Name
	Method  Method `xml:",any"`
}

type Method struct {
	XMLName xml.Name
	Params  []byte `xml:",innerxml"`
}

type ping struct {
}

type register struct {
	UserID string `xml:"user_id"`
	Nama   string `xml:"nama"`
	IP     string `xml:"ip_domisili"`
}

type getSaldo struct {
	UserID string `xml:"user_id"`
}

type transfer struct {
	UserID string `xml:"user_id"`
	Nilai  string `xml:"nilai"`
}

func ewalletServer(c *iris.Context) {
	q := new(RPC)
	c.ReadXML(q)
	var data = `<soapenv:Envelope
		xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
    <soapenv:Header/>
    <soapenv:Body>
        <hy:HelloResponse>
		</hy:HelloResponse>
    	</soapenv:Body>
	</soapenv:Envelope>`
	c.SetContentType("text/xml")
	c.SetBodyString(data)
}
