package main

import (
	"github.com/kataras/go-template/html"
	"gopkg.in/iris-contrib/middleware.v4/logger"
	iris "gopkg.in/kataras/iris.v4"
)

func main() {
	startIris()
}

func startIris() {
	iris.Config.IsDevelopment = true
	iris.UseTemplate(html.New()).Directory("resources/templates", ".html")
	iris.Use(logger.New())
	iris.Get("/", hw2)
	iris.Get("/tugas3/klien", hw3clientget)
	iris.Post("/tugas3/klien", hw3clientpost)
	iris.Post("/tugas3/server", hw3server)
	iris.Get("/tugas3/speksaya.wsdl", hw3spek)
	iris.Get("/tugas4/klien/viewImage/:name", hw4klienviewimage)
	iris.Get("/tugas4/klien/uploadImage", hw4getuploadimage)
	iris.Post("/tugas4/klien/uploadImage", hw4postuploadimage)
	iris.Post("/tugas4/server/postImage", hw4serveruploadimage)
	iris.Get("/tugas4/server/getImage/:name", hw4servergetimage)
	/*
		https://URL_EWALLET_PESERTA_SISDIS/ewallet/ping
		https://URL_EWALLET_PESERTA_SISDIS/ewallet/register
		https://URL_EWALLET_PESERTA_SISDIS/ewallet/getSaldo
		https://URL_EWALLET_PESERTA_SISDIS/ewallet/getTotalSaldo
		https://URL_EWALLET_PESERTA_SISDIS/ewallet/transfer
	*/
	iris.Get("/ewallet/ping", ewalletPing)
	iris.Get("/ewallet/getSaldo", ewalletGetSaldo)
	iris.Post("/ewallet/register", ewalletRegister)
	iris.Listen(":7070")
}
