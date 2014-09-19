package main

import (
	"log"
	"net/http"
	"os"

	"github.com/googollee/go-socket.io"
)

func main() {
	host := ""
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	uri := host + ":" + port

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.Emit("connectied", "hgoehgoe")
		so.On("event", func(msg string) {
			log.Println("emit:", so.Emit("event", msg + "111"))
			so.BroadcastTo("chat", "event", msg + "222")
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	log.Println("Serving at " + uri)
	log.Fatal(http.ListenAndServe(uri, nil))
}
