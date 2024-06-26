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
	"github.com/mrcunninghamz/tprkpr/web/api/empty"
	"github.com/mrcunninghamz/tprkpr/web/api/paydays"
	"github.com/mrcunninghamz/tprkpr/web/api/worksheets"
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
	paydayService := services.NewPaydayService(dataContext.DB)
	billService := services.NewBillService(dataContext.DB)
	worksheetService := services.NewWorksheetService(dataContext.DB)

	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/", index.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/logout", logout.Handler)

	router.GET("/api/payday/edit/:id", paydays.EditForm(paydayService))
	router.GET("/api/paydays", paydays.Get(paydayService))
	router.GET("/api/payday/:id", paydays.View(paydayService))
	router.GET("/api/payday/new", paydays.NewForm)
	router.POST("/api/payday/:id", paydays.Update(paydayService))
	router.POST("/api/payday", paydays.Create(paydayService))
	router.DELETE("/api/payday/:id", paydays.Delete(paydayService))

	router.GET("/api/bill/edit/:id", bills.EditForm(billService))
	router.GET("/api/bill/:id", bills.View(billService))
	router.GET("/api/bill/new/:paydayId", bills.NewForm)
	router.POST("/api/bill/:id", bills.Update(billService))
	router.POST("/api/bill/payday/:paydayId", bills.Create(billService))
	router.DELETE("/api/bill/:id", bills.Delete(billService))

	router.GET("/api/worksheets", worksheets.Get(worksheetService))
	router.GET("/api/worksheet/:id", worksheets.View(worksheetService))
	router.DELETE("/api/worksheet/:id", worksheets.Delete(worksheetService))

	router.POST("/api/worksheet-item/:id", worksheets.UpdateWorksheetItem(worksheetService))

	router.DELETE("/api/empty", empty.New)

	router.StaticFile("/htmx.ext.shoelace.js", "web/htmx.ext.shoelace.js")

	//router.GET("/user", user.Handler)

	return router
}
