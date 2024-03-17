package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
	"github.com/mrcunninghamz/tprkpr/web/app/callback"
	"github.com/mrcunninghamz/tprkpr/web/app/index"
	"github.com/mrcunninghamz/tprkpr/web/app/login"
	"github.com/mrcunninghamz/tprkpr/web/app/logout"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")

	router.GET("/", index.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)
	//router.GET("/user", user.Handler)

	return router
}
