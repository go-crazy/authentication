package routers

import (
	Gin "github.com/gin-gonic/gin"
	"github.com/go-crazy/authentication/controllers"
	"github.com/go-crazy/authentication/core/authentication"
	"github.com/go-crazy/authentication/settings"
)

func SetAuthenticationRoutes(r *Gin.RouterGroup) {

	settings.Init()

	r.POST("/token-auth", controllers.Login)
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(authentication.RequireTokenAuthentication())
	{
		authorized.GET("/refresh-token-auth", controllers.RefreshToken)
		authorized.GET("/logout", controllers.Logout)
	}
	
}
