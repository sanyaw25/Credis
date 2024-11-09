// platform/router/router.go

package router

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"Credis/platform/authenticator"
	"Credis/web/app/attestations"
	"Credis/web/app/callback"
	"Credis/web/app/login"
	"Credis/web/app/logout"
	"Credis/web/app/upload"
	"Credis/web/app/user"
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
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	router.GET("/login_choice", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login_choice.html", nil)
	})

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", user.Handler)
	router.GET("/logout", logout.Handler)
	router.POST("/upload", upload.UploadFile)

	// Create attestation endpoint (requires file URL reference)
	router.POST("/attestations", attestations.CreateAttestation)
	return router
}
