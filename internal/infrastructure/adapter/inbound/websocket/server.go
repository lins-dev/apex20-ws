package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type WSServer struct {
	upgrader websocket.Upgrader
}

func NewWSServer() *WSServer {
	return &WSServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Simplificado para o scaffold inicial
			},
		},
	}
}

func (s *WSServer) GetHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", s.handleConnection)
	return mux
}

func (s *WSServer) Start(port string) error {
	fmt.Printf("WS Service starting on :%s...\n", port)
	return http.ListenAndServe(":"+port, s.GetHandler())
}

func (s *WSServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// Mantém a conexão aberta para o teste
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}
