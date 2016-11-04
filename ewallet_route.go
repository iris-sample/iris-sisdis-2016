package main

import (
	"database/sql"

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

func ewalletPing(c *iris.Context) {
	c.JSON(iris.StatusOK, iris.Map{"pong": 1})
}

func ewalletGetSaldo(c *iris.Context) {
	userID := c.URLParam("user_id")
	var nilaiSaldo int
	db, err := sql.Open("mysql", "root:root@/ewallet")
	stmtOut, err := db.Prepare("SELECT saldo FROM data_pengguna WHERE id = ?")
	err = stmtOut.QueryRow(userID).Scan(&nilaiSaldo)
	if err != nil {
		nilaiSaldo = -1
	}
	db.Close()
	c.JSON(iris.StatusOK, iris.Map{"nilai_saldo": nilaiSaldo})
}

func ewalletGetTotalSaldo(c *iris.Context) {
	// userID := c.URLParam("user_id")
	// Not Found
	// c.JSON(iris.StatusOK, iris.Map{"nilai_saldo": -1})
	c.JSON(iris.StatusOK, iris.Map{"nilai_saldo": 1000})
}

func ewalletRegister(c *iris.Context) {
	req := new(JSONRegister)
	c.ReadJSON(req)
	db, err := sql.Open("mysql", "root:root@/ewallet")
	stmt, _ := db.Prepare("INSERT data_pengguna SET id=?,ip=?,nama=?,saldo=?")
	_, err = stmt.Exec(req.UserID, req.IPDomisili, req.Nama, 100000)
	db.Close()
	if err != nil {
		c.JSON(iris.StatusOK, iris.Map{"error": "Already Registered"})
	}
}

func ewalletTransfer(c *iris.Context) {
	// Fail
	// c.JSON(iris.StatusOK, iris.Map{"status_transfer": -1})
	c.JSON(iris.StatusOK, iris.Map{"status_transfer": 0})
}
