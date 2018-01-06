package jwtauth

import (
	Gin "github.com/gin-gonic/gin"
	"github.com/go-crazy/authentication/routers"
	"github.com/go-crazy/authentication/settings"
)

func Init(r *Gin.RouterGroup) {
	
	settings.Init()
	routers.SetAuthenticationRoutes(r)
}
