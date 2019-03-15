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
        //将username存储到session中
        session := sessions.Default(c)
	session.Set("user", json.Loginname)
	err := session.Save()
	if err != nil {
            code = e.ERROR_AUTH_TOKEN
	} 

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
            "user": json.Loginname,
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

//用户登出
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		//log.Println(user)
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}

//用户信息
func Userinfo(c *gin.Context) {
    //获取session中的user信息
    session := sessions.Default(c)
    user := session.Get("user")
    code := e.INVALID_PARAMS
    if user == nil {
        code = e.ERROR_AUTH_SESSION
    } else {
        code = e.SUCCESS

    }
    //获取GET中token参数
    token := c.Request.URL.Query().Get("token")

    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "msg": e.GetMsg(code),
        "roles": "['admin']",
        "name": user,
        "introduction": "我是超级管理员",
        "token": token,
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
