package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade", err)
		return
	}

	i := 0
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:{
			conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			fmt.Print("ウホウホ","\n")
			i += 1
			}
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	r.Run(":443")
	fmt.Println("listen on 443 port")
}
