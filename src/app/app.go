package app

import (
	"github.com/dmolina79/bookstore_oauth-api/src/domain/access_token"
	"github.com/dmolina79/bookstore_oauth-api/src/http"
	"github.com/dmolina79/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.New())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")

}
