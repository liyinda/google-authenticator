package apis

import (
    "github.com/gin-gonic/gin"
    "github.com/liyinda/google-authenticator/api/models"
    "github.com/gin-gonic/contrib/sessions"
    "net/http"
    //"fmt"
    //"strings"
    //"github.com/liyinda/google-authenticator/pkg/util"
    "github.com/liyinda/google-authenticator/pkg/e"
)


//用户信息
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
    json.Qrcode = "text111"

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

