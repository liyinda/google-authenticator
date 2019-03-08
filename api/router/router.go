package router

import (
    "github.com/gin-gonic/gin"
    . "github.com/liyinda/google-authenticator/api/apis"
    "net/http"
    "fmt"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    //登录入口
    passport := router.Group("/passport")
    {
        passport.POST("/login", Login) 
        //passport.GET("/userinfo", userinfo) 
    }

    //用户管理入口
    home := router.Group("/home")
    {
        home.GET("/userinfo", Userinfo) 
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
