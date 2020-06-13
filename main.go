package main

import (
	"github.com/fatih/color"
	mysql "github.com/koushamad/goro/app/Model/Mysql"
	"github.com/koushamad/goro/app/Model/mongo"
	"github.com/koushamad/goro/app/socket"
	"log"
	"net/http"
)

func main() {
	Boot()
	//setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle("/io/", socket.Init())
	color.Green("start to listen port 9000")
	log.Panic(http.ListenAndServe(":9000", serveMux))
}

func Boot() {
	socket.Init().Handler()
	mysql.Init()
	mongo.Init()
}
