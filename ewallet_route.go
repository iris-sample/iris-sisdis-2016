package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	iris "gopkg.in/kataras/iris.v4"
)

type JSONRegister struct {
	UserID     string `json:"user_id"`
	Nama       string `json:"nama"`
	IPDomisili string `json:"ip_domisili"`
}

type JSONTransfer struct {
	UserID string `json:"user_id"`
	Nilai  int    `json:"nilai"`
}

type JSONPing struct {
	Pong int `json:"pong"`
}

func ewalletPing(c *iris.Context) {
	c.JSON(iris.StatusOK, iris.Map{"pong": 1})
}

func ewalletGetSaldo(c *iris.Context) {
	if checkQuorum() >= 5 {
		userID := c.URLParam("user_id")
		var nilaiSaldo int
		db, err := sql.Open("mysql", "root:password2016@/ewallet")
		stmtOut, err := db.Prepare("SELECT saldo FROM data_pengguna WHERE id = ?")
		err = stmtOut.QueryRow(userID).Scan(&nilaiSaldo)
		if err != nil {
			nilaiSaldo = -1
		}
		db.Close()
		c.JSON(iris.StatusOK, iris.Map{"nilai_saldo": nilaiSaldo})
	} else {
		c.JSON(iris.StatusOK, iris.Map{"error": "Quorum tidak tercapai"})
	}
}

func ewalletGetTotalSaldo(c *iris.Context) {
	if checkQuorum() == 8 {
		// userID := c.URLParam("user_id")
		// Not Found
		// c.JSON(iris.StatusOK, iris.Map{"nilai_saldo": -1})
		c.JSON(iris.StatusOK, iris.Map{"nilai_saldo": 1000})
	} else {
		c.JSON(iris.StatusOK, iris.Map{"error": "Quorum tidak tercapai"})
	}
}

func ewalletRegister(c *iris.Context) {
	if checkQuorum() >= 5 {
		req := new(JSONRegister)
		c.ReadJSON(req)
		db, err := sql.Open("mysql", "root:password2016@/ewallet")
		stmt, _ := db.Prepare("INSERT data_pengguna SET id=?,ip=?,nama=?,saldo=?")
		_, err = stmt.Exec(req.UserID, req.IPDomisili, req.Nama, 1000000)
		db.Close()
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"error": "Already Registered"})
		}
	} else {
		c.JSON(iris.StatusOK, iris.Map{"error": "Quorum tidak tercapai"})
	}
}

func ewalletTransfer(c *iris.Context) {
	if checkQuorum() >= 5 {
		req := new(JSONTransfer)
		c.ReadJSON(req)
		// Fail
		// c.JSON(iris.StatusOK, iris.Map{"status_transfer": -1})
		c.JSON(iris.StatusOK, iris.Map{"status_transfer": 0})
	} else {
		c.JSON(iris.StatusOK, iris.Map{"error": "Quorum tidak tercapai"})
	}
}

func checkQuorum() (total int) {
	total = 0
	listIP := [8]string{"192.168.75.70", "192.168.75.75", "192.168.75.77", "192.168.75.78", "192.168.75.81", "192.168.75.93", "192.168.75.98", "192.168.75.105"}
	for x := 0; x < 8; x++ {
		request, err := http.Get("https://" + listIP[x] + "/ewallet/ping")
		if err != nil {
		} else {
			decodeJSON := json.NewDecoder(request.Body)
			var data JSONPing
			err = decodeJSON.Decode(&data)
			if err != nil {
			} else {
				total += data.Pong
			}
		}
	}
	return
}
