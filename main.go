package main

import (
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
    "io"
    "golang.org/x/net/websocket"
)

func EchoHandler(ws *websocket.Conn) {
    io.Copy(ws, ws)
}

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())

	grpc.NewServer()
}
