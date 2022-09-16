package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liyinda/google-authenticator/auth"
	"github.com/liyinda/google-authenticator/home"
	"github.com/liyinda/google-authenticator/passport"
	"github.com/liyinda/google-authenticator/utils/rbac"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//前端跨域问题
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		//AllowHeaders:     []string{"Authorization", "Content-Length", "Content-Type","X-Token", "Access-Control-Allow-Origin"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "X-Token", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		AllowOrigins:     []string{"http://localhost:7000"},
	}))

	//引用静态资源
	r.LoadHTMLGlob("dist/*.html")
	//r.LoadHTMLFiles("static/*/*")
	r.Static("/static", "./dist/static")
	r.StaticFile("/vue/", "dist/index.html")

	//设置sessions
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	p := r.Group("/passport")
	p.Use()
	{
		p.POST("/login", passport.Login)
		p.POST("/logout", passport.Logout)
	}

	d := r.Group("/auth")
	{
		d.GET("/api/v1", auth.ApiQrcode)
	}

	h := r.Group("/home")
	h.Use(rbac.JWT())
	//h.Use(rbac.JWTAuthMiddleWare())
	{
		h.GET("/userinfo", home.Userinfo)
		h.POST("/create", home.AuthUserCreate)
		h.GET("/list", home.AuthUserList)
		h.GET("/qrcode", home.GetQrcode)
		h.POST("/delete", home.AuthUserDelete)
	}
	h.Use(rbac.AuthRequired())

	return r
}
