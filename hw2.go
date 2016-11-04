package main

import (
	"database/sql"
	"os/exec"

	iris "gopkg.in/kataras/iris.v4"

	_ "github.com/go-sql-driver/mysql"
)

func hw2(c *iris.Context) {
	db, err := sql.Open("mysql", "root:password2016@/sisdis")
	var hello string
	var uptime string
	cmdOut, err := exec.Command("uptime").Output()
	uptime = string(cmdOut)
	stmt, err := db.Prepare("UPDATE hw2 SET info = ? WHERE id = ?")
	stmt.Exec(uptime, 2)
	stmtOut, err := db.Prepare("SELECT info FROM hw2 WHERE id = ?")
	err = stmtOut.QueryRow(1).Scan(&hello)
	stmtOut, err = db.Prepare("SELECT info FROM hw2 WHERE id = ?")
	err = stmtOut.QueryRow(2).Scan(&uptime)
	_ = err
	db.Close()
	c.MustRender("hw2.html", struct {
		Hello  string
		Uptime string
	}{Hello: hello, Uptime: uptime})
}
