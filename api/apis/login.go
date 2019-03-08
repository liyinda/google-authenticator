package apis

import (
    "github.com/gin-gonic/gin"
    models "github.com/liyinda/google-authenticator/api/models"
    "github.com/gin-gonic/contrib/sessions"
    "net/http"
)

//用户登录
func Login(c *gin.Context) {
     var json models.LoginJson
     if err := c.ShouldBindJSON(&json); err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
     }
     if json.Loginname == "liyinda" && json.Password == "123456" {
         c.JSON(http.StatusOK, gin.H{
             "status": 200,
             "token": "sdfsdf",
             "message": "you are logged in",
         })
         return
     } else {
             c.JSON(401, gin.H{"status": "unauthorized"})
     }
}

//用户信息
func Userinfo(c *gin.Context) {
    c.JSON(200, gin.H{
        "roles": "['admin']",
        "name": "Super Admin",
        "introduction": "我是超级管理员",
        "token": "admin",
        "avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
    })

}

//用户会话保持
func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        if user == nil {
            // You'd normally redirect to login page
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
        } else {
            // Continue down the chain to handler etc
            c.Next()
        }
    }
}
