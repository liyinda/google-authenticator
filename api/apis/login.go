package apis

import (
    "github.com/gin-gonic/gin"
    models "github.com/liyinda/google-authenticator/api/models"
    "github.com/gin-gonic/contrib/sessions"
    "net/http"
    //"fmt"
    "crypto/md5"
    "encoding/hex"
    "strings"
    //"strconv"
    //simplejson "github.com/bitly/go-simplejson"
    //"reflect"
)

//用户登录
func Login(c *gin.Context) {
    var json models.LoginJson
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    //获取管理员密码
    var authcms_admin models.Authcms_admin
    result, err := authcms_admin.GetPassword(json.Loginname)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "message": "program error",
        })
        return
    }
    if len(result) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "unknown user",
        })
        return
    }
    //获取Authcms_admin结构体
    admin := result[0]

    //密码加盐
    salt := md5.New() 
    salt.Write([]byte(admin.Salt + json.Password))
    cipherStr := salt.Sum(nil)
    //转换成大写字符
    saltPassword := strings.ToUpper(hex.EncodeToString(cipherStr))

    //比较用户POST提交的密码是否与数据库中密码一致
    if saltPassword == admin.Password {
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

////管理员
//func UserList(c *gin.Context) {
//    var user models.Userlist
//    //user.Username = c.Request.FormValue("username")
//    //user.Password = c.Request.FormValue("password")
//    result, err := user.Userlist(2)
//    //js := result[1]
//    //js, err := simplejson.NewJson([]byte(result))
//    
//
//    fmt.Println("type:", reflect.TypeOf(result))
//    shu := result[0]
//    fmt.Println(shu.Password)
//
//    if err != nil {
//        c.JSON(http.StatusOK, gin.H{
//            "code":    -1,
//            "message": "抱歉未找到相关信息",
//        })
//        return
//    }
//
//    c.JSON(http.StatusOK, gin.H{
//        "code": 1,
//        "data":   result,
//    })
//}

//func Update(c *gin.Context) {
//    var user models.User
//    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
//    user.Password = c.Request.FormValue("password")
//    result, err := user.Update(id)
//    if err != nil || result.ID == 0 {
//        c.JSON(http.StatusOK, gin.H{
//            "code":    -1,
//            "message": "修改失败",
//        })
//        return
//    }
//    c.JSON(http.StatusOK, gin.H{
//        "code":  1,
//        "message": "修改成功",
//    })
//}


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
