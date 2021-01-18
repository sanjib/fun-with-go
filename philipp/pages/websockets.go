package pages

import "net/http"

func WebSocketsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/websockets.html")
}
