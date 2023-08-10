package database

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var SessionStore *session.Store

func InitSessionStore() {
	store := redis.New(redis.Config{
		Host:     "redis",
		Port:     6379,
		Database: 0,   
		Password: "",   
	})

	SessionStore = session.New(session.Config{
		Storage:       store,
		//swap value to true before production
		CookieHTTPOnly: false,             // Set the cookie as httpOnly
		CookieSecure:   false,             // Set the cookie to be secure (only transmitted over HTTPS)
		//CookieSameSite: "Strict",         // Optional: set the SameSite attribute for the cookie. It can prevent CSRF attacks in some cases
		Expiration:     1 *time.Hour ,              // Expire after 24 hours
	})

	fmt.Println("Session Started!")
}
