package main

import iris "gopkg.in/kataras/iris.v4"

type ReplyHello struct {
	HelloResponse string `xml:"Body>HelloResponse"`
}

func hw3server(c *iris.Context) {
	q := new(RequestHello)
	c.ReadXML(q)
	var data = `<soapenv:Envelope
		xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:hy="http://www.herongyang.com/Service/">
    <soapenv:Header/>
    <soapenv:Body>
        <hy:HelloResponse>`
	data += "Hello " + q.HelloRequest + "! This message is from Prakash :D"
	data += `</hy:HelloResponse>
    	</soapenv:Body>
	</soapenv:Envelope>`
	c.SetContentType("text/xml")
	c.SetBodyString(data)
}
