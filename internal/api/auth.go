package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	gsession "github.com/gorilla/sessions"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/util"
	"net/http"
	"time"
)

var store = gsession.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))

type Users struct {
	CreatedAt time.Time
	Password  string `json:"password", db:"password"`
	Username  string `json:"username", db:"username"`
	UUID      string
}

// POST /api/v1/signup
func Signup(router *gin.RouterGroup, conf *config.Config) {
	router.POST("/signup", func(c *gin.Context) {
		params := Users{}
		userMng, _ := photoprism.NewUserManager(*conf)

		if err := c.BindJSON(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.UcFirst(err.Error())})
			return
		}

		if params.Username == "" || params.Password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Missing username or password")})
		}

		if userMng.HasUser(params.Username) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("\"%s\" already exists", &params)})
			return
		}

		user := userMng.AddUser(params.Username, params.Password)
		jsontoken := photoprism.GetJSONToken(user)

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write([]byte(jsontoken))
	})
}

// POST /api/v1/signin
func Signin(router *gin.RouterGroup, conf *config.Config) {
	router.POST("/signin", func(c *gin.Context) {
		userMng, _ := photoprism.NewUserManager(*conf)
		params := Users{}
		if err := c.BindJSON(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": util.UcFirst(err.Error())})
			return
		}

		if params.Username == "" || params.Password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Missing username or password")})
		}

		user := userMng.FindUser(params.Username)
		if user.Username == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("\"%s\" Username or password was wrong")})
			return
		}

		if !userMng.CheckPassword(user.Password, params.Password) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("\"%s\" Username or password was wrong")})
			return
		}

		jsontoken := photoprism.GetJSONToken(user)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Write([]byte(jsontoken))
	})
}
