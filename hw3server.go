package main

import (
	"encoding/xml"

	"github.com/kataras/iris"
)

type ReplyHello struct {
	XMLName       xml.Name `xml:"Envelope"`
	Header        string   `xml:"Header"`
	HelloResponse string   `xml:"Body>HelloResponse"`
}

func hw3server(c *iris.Context) {
	q := new(RequestHello)
	c.ReadXML(q)
	x := new(ReplyHello)
	x.HelloResponse = "Hello " + q.HelloRequest + "! This message is from Prakash :D"
	c.XML(iris.StatusOK, x)
}
