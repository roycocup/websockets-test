package main

import (
	"net/http"

	"./lib"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		_, err := upgrader.Upgrade(w, r, nil)
		lib.Check(err)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		spew.Dump(r.Form.Encode())
	})

	logrus.Info("Server started...")
	http.ListenAndServe(":3000", nil)
}
