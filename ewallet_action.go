package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	iris "gopkg.in/kataras/iris.v4"
)

func ewalletAction(c *iris.Context) {
	action := c.URLParam("action")
	if action == "ping" {
		request, err := http.Get("https://" + c.URLParam("addr") + "/ewallet/ping")
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error"})
			return
		}
		decodeJSON := json.NewDecoder(request.Body)
		var data JSONPing
		err = decodeJSON.Decode(&data)
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error"})
			return
		}
		c.JSON(iris.StatusOK, iris.Map{"status": "ok"})
	} else if action == "register" {
		id := c.URLParam("id")
		nama := c.URLParam("nama")
		ip := c.URLParam("addr")
		var payload JSONRegister
		payload.UserID = id
		payload.Nama = nama
		payload.IPDomisili = ip
		out, err := json.Marshal(payload)
		request, err := http.Post("https://prakash.sisdis.ui.ac.id/eewallet/register", "application/json", bytes.NewBuffer(out))
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Pastikan anda terhubung internet."})
			return
		}
		decodeJSON := json.NewDecoder(request.Body)
		var data JSONReplyRegister
		_ = decodeJSON.Decode(&data)
		if data.Status == "error" {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": data.Reason})
		} else {
			c.JSON(iris.StatusOK, iris.Map{"status": "ok"})
		}
	} else if action == "transfer" {

	} else if action == "getSaldo" {

	} else if action == "getTotalSaldo" {

	}
}
