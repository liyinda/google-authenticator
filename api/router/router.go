package router

import (
    "github.com/gin-gonic/gin"
    . "github.com/liyinda/google-authenticator/api/apis"
    "net/http"
    "fmt"
    "github.com/liyinda/google-authenticator/middleware/jwt"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-contrib/cors"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    //引用静态资源
    router.LoadHTMLGlob("dist/*.html")
    router.LoadHTMLFiles("static/*/*")
    router.Static("/static", "./dist/static")
    router.StaticFile("/vue/", "dist/index.html")

    //设置sessions
    store := sessions.NewCookieStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))

    //登录入口
    passport := router.Group("/passport")
    {
        passport.POST("/login", Login) 
    }
    //passport.Use(AuthRequired())

    passport.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        //AllowOrigins:     []string{"http://101.200.42.56:8888"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
        AllowHeaders:     []string{"Content-Type,Authorization,X-Token"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "*"
            //return origin == "http://101.200.42.56:8888"
        },
    }))



    //用户管理入口
    home := router.Group("/home")
    home.Use(jwt.JWT())
    {
        home.GET("/userinfo", Userinfo) 
        home.POST("/useradd", Useradd) 
    }
    home.Use(AuthRequired())

    //定义默认路由
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": 404,
            "error":  "404, page not exists!",
        })
    })


    //设置cookie
    router.GET("/cookie", func(c *gin.Context) {

        cookie, err := c.Cookie("gin_cookie")

        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }

        fmt.Printf("Cookie value: %s \n", cookie)
    })

    return router
}
