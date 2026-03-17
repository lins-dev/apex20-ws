package port

import "net/http"

// WSServer define a interface para o servidor de WebSockets.
type WSServer interface {
	Start(port string) error
	GetHandler() http.Handler
}
