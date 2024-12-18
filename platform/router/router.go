package router

import (
	"encoding/gob"
	"io/ioutil"
	"net/http"
	"os"

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

func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

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

	router.GET("/check-output", func(c *gin.Context) {
		if _, err := os.Stat("/mnt/Disk_2/hack_cbs/project/Credis/output.txt"); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"status": "File not yet ready"})
			return
		}

		data, err := ioutil.ReadFile("/mnt/Disk_2/hack_cbs/project/Credis/output.txt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "File ready", "content": string(data)})
	})

	router.POST("/attestations", attestations.CreateAttestation)
	return router
}
