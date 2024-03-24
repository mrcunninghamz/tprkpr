package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/mrcunninghamz/tprkpr/platform/authenticator"
	"github.com/mrcunninghamz/tprkpr/platform/data"
	"github.com/mrcunninghamz/tprkpr/platform/services"
	"github.com/mrcunninghamz/tprkpr/web/api/bills" // assuming this is the package name
	"github.com/mrcunninghamz/tprkpr/web/api/paydays"
	"github.com/mrcunninghamz/tprkpr/web/app/callback"
	"github.com/mrcunninghamz/tprkpr/web/app/index"
	"github.com/mrcunninghamz/tprkpr/web/app/login"
	"github.com/mrcunninghamz/tprkpr/web/app/logout"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator, dataContext *data.DataContext) *gin.Engine {
	router := gin.Default()
	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	paydayService := services.NewPayDayService(dataContext.DB)
	billService := services.NewBillService(dataContext.DB)

	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/", index.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)
	router.GET("/api/paydays", paydays.Get(paydayService))

	router.GET("/views/bill/edit/:id", bills.Edit(billService))
	router.POST("/views/bill/edit/:id", bills.Update(billService))
	router.GET("/views/bill/:id", bills.View(billService))

	router.StaticFile("/htmx.ext.shoelace.js", "web/htmx.ext.shoelace.js")

	//router.GET("/user", user.Handler)

	return router
}
