package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	iris "gopkg.in/kataras/iris.v4"
)

type JSONReply struct {
	Status         string `json:"status"`
	Reason         string `json:"reason"`
	NilaiSaldo     int    `json:"nilai_saldo"`
	StatusTransfer int    `json:"status_transfer"`
}

func ewalletAction(c *iris.Context) {
	action := c.URLParam("method")
	if action == "ping" {
		ip := c.URLParam("addr")
		request, err := http.Get("https://" + ip + "/ewallet/ping")
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
		request, err := http.Post("https://prakash.sisdis.ui.ac.id/ewallet/register", "application/json", bytes.NewBuffer(out))
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Pastikan anda terhubung internet."})
			return
		}
		decodeJSON := json.NewDecoder(request.Body)
		var data JSONReply
		_ = decodeJSON.Decode(&data)
		if data.Status == "error" {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": data.Reason})
		} else {
			c.JSON(iris.StatusOK, iris.Map{"status": "ok"})
		}
	} else if action == "transfer" {
		id := c.URLParam("id")
		request, err := http.Get("https://prakash.sisdis.ui.ac.id/ewallet/getSaldo?user_id=" + id)
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Pastikan anda terhubung internet."})
			return
		}
		decodeJSON := json.NewDecoder(request.Body)
		var data JSONReply
		_ = decodeJSON.Decode(&data)
		if data.Status == "error" {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": data.Reason})
		} else {
			oldSaldo := data.NilaiSaldo
			if oldSaldo == -1 {
				c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "User tidak terdaftar."})
			} else {
				nilai, _ := strconv.Atoi(c.URLParam("nilai"))
				if oldSaldo < nilai {
					c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Saldo tidak cukup."})
				} else {
					ip := c.URLParam("addr")
					var payload JSONTransfer
					payload.UserID = id
					payload.Nilai = nilai
					out, err := json.Marshal(payload)
					request, err := http.Post("https://"+ip+"/ewallet/transfer", "application/json", bytes.NewBuffer(out))
					if err != nil {
						c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Pastikan anda terhubung internet."})
						return
					}
					decodeJSON := json.NewDecoder(request.Body)
					var data JSONReply
					_ = decodeJSON.Decode(&data)
					if data.Status == "error" {
						c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": data.Reason})
					} else {
						if data.StatusTransfer == -1 {
							c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "User tidak terdaftar."})
						} else {
							var nilaiSaldo int
							db, _ := sql.Open("mysql", "root:password2016@/ewallet")
							stmtOut, _ := db.Prepare("SELECT saldo FROM data_pengguna WHERE id = ?")
							err = stmtOut.QueryRow(id).Scan(&nilaiSaldo)
							newSaldo := nilaiSaldo - nilai
							stmt, _ := db.Prepare("UPDATE data_pengguna SET saldo = ? WHERE id = ? ")
							_, err = stmt.Exec(newSaldo, id)
							db.Close()
							c.JSON(iris.StatusOK, iris.Map{"status": "ok"})
						}
					}
				}
			}
		}
	} else if action == "getSaldo" {
		id := c.URLParam("id")
		request, err := http.Get("https://prakash.sisdis.ui.ac.id/ewallet/getSaldo?user_id=" + id)
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Pastikan anda terhubung internet."})
			return
		}
		decodeJSON := json.NewDecoder(request.Body)
		var data JSONReply
		_ = decodeJSON.Decode(&data)
		if data.Status == "error" {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": data.Reason})
		} else {
			if data.NilaiSaldo == -1 {
				c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "User tidak terdaftar."})
			} else {
				c.JSON(iris.StatusOK, iris.Map{"status": "ok", "nilai_saldo": data.NilaiSaldo})
			}
		}
	} else if action == "getTotalSaldo" {
		id := c.URLParam("id")
		request, err := http.Get("https://prakash.sisdis.ui.ac.id/ewallet/getTotalSaldo?user_id=" + id)
		if err != nil {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "Pastikan anda terhubung internet."})
			return
		}
		decodeJSON := json.NewDecoder(request.Body)
		var data JSONReply
		_ = decodeJSON.Decode(&data)
		if data.Status == "error" {
			c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": data.Reason})
		} else {
			if data.NilaiSaldo == -1 {
				c.JSON(iris.StatusOK, iris.Map{"status": "error", "reason": "User tidak terdaftar."})
			} else {
				c.JSON(iris.StatusOK, iris.Map{"status": "ok", "nilai_saldo": data.NilaiSaldo})
			}
		}
	}
}
