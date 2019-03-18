package apis

import (
    "github.com/gin-gonic/gin"
    "github.com/liyinda/google-authenticator/api/models"
    "github.com/gin-gonic/contrib/sessions"
    "net/http"
    //"fmt"
    //"strings"
    "github.com/liyinda/google-authenticator/pkg/util"
    "github.com/liyinda/google-authenticator/pkg/e"
    "strconv"
)


//新增用户信息
func Useradd(c *gin.Context) {
    //获取session中的user信息
    session := sessions.Default(c)
    user := session.Get("user")
    code := e.INVALID_PARAMS
    if user == nil {
        code = e.ERROR_AUTH_SESSION
    } else {
        code = e.SUCCESS

    }
    //获取POST中json参数
    var json models.Authcms_user

    if err := c.ShouldBindJSON(&json); err != nil {
        code = e.ERROR_NOT_JSON
    }
    /*
    fmt.Println(json.ID)
    fmt.Println(json.User_name)
    fmt.Println(json.Phone)
    fmt.Println(json.Qrcode)
    */

    //根据提交用户名称生成二维码秘钥和图片base64编码格式
    //json.Secretid, json.Qrcode, err := util.CreateQrcode(json.User_name) 
    secretid, qrcode, err := util.CreateQrcode(json.User_name)     
    if err != nil {
        code = e.ERROR
    }
    json.Secretid = secretid
    json.Qrcode = qrcode

    id, err := json.Useradd()
    if err != nil {
        code = e.ERROR_NOT_JSON
    }


    token := c.Request.URL.Query().Get("token")

    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "msg": e.GetMsg(code),
        "name": user,
        "token": token,
        "id": id,
    })
}

//更改用户信息
func Useredit(c *gin.Context) {
    //获取session中的user信息
    session := sessions.Default(c)
    user := session.Get("user")
    code := e.INVALID_PARAMS
    if user == nil {
        code = e.ERROR_AUTH_SESSION
    } else {
        code = e.SUCCESS

    }
    //获取POST中json参数
    var json models.Authcms_user

    if err := c.ShouldBindJSON(&json); err != nil {
        code = e.ERROR_NOT_JSON
    }
    /*
    fmt.Println(json.ID)
    fmt.Println(json.User_name)
    fmt.Println(json.Phone)
    */
    //更新用户信息表
    //id, err := strconv.ParseInt(json.ID, 10, 64)
    result, err := json.Useredit(json.ID)
    if err != nil || result.ID == 0 {
        code = e.ERROR
    } else {
        code = e.SUCCESS
    }

    token := c.Request.URL.Query().Get("token")

    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "msg": e.GetMsg(code),
        "name": user,
        "token": token,
        "id": json.ID,
    })
}

//更改用户信息
func Userlist(c *gin.Context) {
    //获取session中的user信息
    session := sessions.Default(c)
    user := session.Get("user")
    code := e.INVALID_PARAMS
    if user == nil {
        code = e.ERROR_AUTH_SESSION
    } else {
        code = e.SUCCESS

    }
    //获取POST中json参数
    var json models.Authcms_user

    if err := c.ShouldBindJSON(&json); err != nil {
        code = e.ERROR_NOT_JSON
    }
    //获取url中token, page, limit
    token := c.Request.URL.Query().Get("token")
    page := c.Request.URL.Query().Get("page")
    limit := c.Request.URL.Query().Get("limit")

    pageint, _ := strconv.ParseInt(page, 10, 64)
    limitint, _ := strconv.ParseInt(limit, 10, 64)


    //获取用户信息表
    result, err := json.Userlist(pageint, limitint)
    if err != nil{
        code = e.ERROR
    } else {
        code = e.SUCCESS
    }

    c.JSON(http.StatusOK, gin.H{
        "status": code,
        "msg": e.GetMsg(code),
        "name": user,
        "token": token,
        "id": json.ID,
        "users": result,
    })
}

