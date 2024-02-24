package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ganasa18/document-be/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type WebSocketControllerImpl struct {
}

func NewWebSocketController() WebSocketController {
	return &WebSocketControllerImpl{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	send chan string
}

var clients = make(map[*Client]bool)
var broadcast = make(chan string)

func (controller *WebSocketControllerImpl) HandlerWebSocketController(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	utils.PanicIfError(err)
	defer conn.Close()

	client := &Client{
		conn: conn,
		send: make(chan string),
	}

	clients[client] = true

	go handleClientMessages(client)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients, client)
			close(client.send)
			return
		}

		message := string(p)
		broadcast <- message

		fmt.Println(messageType, "MESSAGE TYPE SERVER")
		log.Printf("Received message: %s\n", p)

	}
}

func handleClientMessages(client *Client) {
	for message := range client.send {

		err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

func (controller *WebSocketControllerImpl) HandleMessages() {
	for message := range broadcast {
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(clients, client)
			}
		}
	}
}
