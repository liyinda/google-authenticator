package apis

import (
    "github.com/gin-gonic/gin"
    "github.com/liyinda/google-authenticator/api/models"
    "net/http"
    "fmt"
    //"strings"
    "github.com/liyinda/google-authenticator/pkg/util"
    //"github.com/liyinda/google-authenticator/pkg/e"
    "strconv"
)


//获取用户名返回6位验证码
func ApiQrcode(c *gin.Context) {
    var json models.Authcms_user
    //获取请求中issuser参数
    issuser := c.Request.URL.Query().Get("issuser")
    code := c.Request.URL.Query().Get("code")

    var flag string
    //查询用户信息
    result, err := json.Usersearch(issuser)
    if err == nil {
        flag = "no user"
    }
    //获取6位验证码
    vcode := util.ReturnCode(result.Secretid)
    //将unit32转换成int类型
    icode := strconv.Itoa(int(vcode))

    if string(icode) == code {
        flag = "ok"
    } else {
        flag = "error"
    }

    c.String(http.StatusOK, fmt.Sprintf("%s", flag))
 
}
