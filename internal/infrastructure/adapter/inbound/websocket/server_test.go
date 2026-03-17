package websocket_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	adapter "github.com/apex20/ws/internal/infrastructure/adapter/inbound/websocket"
)

func TestWSServer_Connection(t *testing.T) {
	// Setup do Servidor (Ainda não existe)
	server := adapter.NewWSServer()
	ts := httptest.NewServer(server.GetHandler())
	defer ts.Close()

	// Convertendo URL para WS
	url := "ws" + strings.TrimPrefix(ts.URL, "http")

	// Tentando conectar
	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial(url+"/ws", nil)
	
	// Asserção
	assert.NoError(t, err)
	defer ws.Close()
}
