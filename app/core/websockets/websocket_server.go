package websockets

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
)

type WebSocketServer interface {
	HandleMessages()
	HandleWebSocket(ctx *websocket.Conn)
	BroadcastCmd(cmd WsCommand)
}

type webSocketServer struct {
	id        string
	clients   map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocket() WebSocketServer {
	return &webSocketServer{
		id:        uuid.New().String(),
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

func (s *webSocketServer) HandleConnections(ctx *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(ctx) {
		return ctx.Next()
	}
	return fiber.ErrUpgradeRequired
}

func (s *webSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	// Register a new Client
	s.clients[ctx] = true
	defer func() {
		delete(s.clients, ctx)
		err := ctx.Close()
		if err != nil {
			log.Printf("WS ctx.Close  Error: %v ", err)
			return
		}
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error:", err)
			break
		}

		// send the message to the broadcast channel
		log.Println(string(msg))
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Fatalf("Error Unmarshalling")
		}
		message.Client = s.id

		s.broadcast <- &message
	}
}

func (s *webSocketServer) HandleMessages() {
	for {
		msg := <-s.broadcast
		// Send the message to all Clients
		for client := range s.clients {
			err := client.WriteJSON(WsCommand{
				Command: msg.Action,
				Code:    200,
				Msg:     msg.Value,
				Data:    nil,
			})
			if err != nil {
				log.Printf("Write  Error: %v ", err)
				err := client.Close()
				if err != nil {
					log.Printf("WS client.Close  Error: %v ", err)
				}
				delete(s.clients, client)
			}
		}
	}
}

func (s *webSocketServer) BroadcastCmd(cmd WsCommand) {
	for client := range s.clients {
		err := client.WriteJSON(cmd)
		if err != nil {
			log.Printf("Write  Error: %v ", err)
			err := client.Close()
			if err != nil {
				log.Printf("WS client.Close  Error: %v ", err)
			}
			delete(s.clients, client)
		}
	}
}
