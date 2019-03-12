package apis

import (
    "github.com/gin-gonic/gin"
    //models "github.com/liyinda/google-authenticator/api/models"
    "github.com/liyinda/google-authenticator/api/models"
    "github.com/gin-gonic/contrib/sessions"
    "net/http"
    //"fmt"
    "strings"
    "github.com/liyinda/google-authenticator/pkg/util"
    "github.com/liyinda/google-authenticator/pkg/e"
)

//用户登录
func Login(c *gin.Context) {
    var json models.LoginJson

    code := e.INVALID_PARAMS
    if err := c.ShouldBindJSON(&json); err != nil {
        code = e.ERROR_NOT_JSON
        c.JSON(http.StatusBadRequest, gin.H{
            "msg": e.GetMsg(code),
        })
        return
    }
    //获取管理员密码
    var authcms_admin models.Authcms_admin
    result, err := authcms_admin.GetPassword(json.Loginname)
    if err != nil {
        code = e.ERROR
        return
    }
    if len(result) == 0 {
        code = e.ERROR_NOT_EXIST_USER
        c.JSON(http.StatusOK, gin.H{
            "msg": e.GetMsg(code),
        })
        return
    }
    //获取Authcms_admin结构体
    admin := result[0]

    //密码加盐
    salt := util.EncodeMD5(admin.Salt + json.Password)
    //转换成大写字符
    saltPassword := strings.ToUpper(salt)

    //比较用户POST提交的密码是否与数据库中密码一致
    if saltPassword == admin.Password {
        code = e.SUCCESS
        //获取用户token信息
        token, err := util.GenerateToken(json.Loginname, json.Password)
        if err != nil {
            code = e.ERROR_AUTH_TOKEN
        } 

        c.JSON(http.StatusOK, gin.H{
            "status": code,
            "token": token,
            "msg": e.GetMsg(code),
        })
        return
    } else {
        code = e.ERROR
    }
    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "msg": e.GetMsg(code),
    })

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
