package pages

import (
	"database/sql"
	"time"

	"github.com/gorilla/sessions"
)

const (
	cookieName           = "viaye-go-ph-3000"
	sessKeyFooBar        = "foobar"
	sessKeyAuthenticated = "authenticated"
)

var (
	// DB provides db connection
	DB *sql.DB

	cookieStoreKey = []byte("oreo-cream-chocolate-chips")
	cookieStore    = sessions.NewCookieStore(cookieStoreKey)
)

type Game struct {
	Name      string    `json:"name"`
	Developer string    `json:"developer"`
	Release   time.Time `json:"release"`
	Genre     string    `json:"genre"`
}
