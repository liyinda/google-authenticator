package home

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liyinda/google-authenticator/store/sqlite"
	"github.com/liyinda/google-authenticator/utils/errno"
	"github.com/liyinda/google-authenticator/utils/handler"
	"github.com/liyinda/google-authenticator/utils/qrcode"
	"log"
	"net/http"
	"strconv"
)

//用户信息
func Userinfo(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")

	code := errno.INVALID_PARAMS
	if user == nil {
		code = errno.ERROR_AUTH_SESSION
	} else {
		code = errno.SUCCESS
	}
	log.Println("code:", code, "session", user)

	//获取GET中token参数
	token := c.Request.URL.Query().Get("token")

	data := gin.H{
		"msg":          errno.GetMsg(code),
		"roles":        "['admin']",
		"name":         user,
		"introduction": "我是超级管理员",
		"token":        token,
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}

func AuthUserCreate(c *gin.Context) {
	var authuser sqlite.Authuser
	//将request的body中的数据，自动按照json格式解析到结构体。
	if err := c.ShouldBindJSON(&authuser); err != nil {
		handler.SendResponse(c, errno.ErrStruct, nil)
		return
	}

	//判断项目名称是否相同，在相同的环境和命名空间下
	count, _ := sqlite.CountAuthUserByName(authuser.Name)
	if count > 0 {
		handler.SendResponse(c, errno.ErrProjectCreate, "error, use the same authuser name.")
		return
	}
	secretid, qrcode, err := qrcode.CreateQrcode(authuser.Name)
	authuser.Secretid = secretid
	authuser.Qrcode = qrcode
	if err != nil {
		handler.SendResponse(c, errno.ErrProjectCreate, nil)
		return
	}

	err = sqlite.CreateAuthUser(authuser)
	if err != nil {
		handler.SendResponse(c, errno.ErrProjectCreate, nil)
		return
	}
	handler.SendResponse(c, errno.OK, "success")
}

func AuthUserList(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		handler.SendResponse(c, errno.ErrToken, "error, session is nil")
	}

	aus, err := sqlite.ListAuthUser()
	if err != nil {
		handler.SendResponse(c, errno.ErrQuery, "error, list auth user!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errno.SUCCESS,
		"data": gin.H{
			"items": aus,
			"total": 2,
		},
	})

}

func GetQrcode(c *gin.Context) {
	//获取session中的user信息
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		handler.SendResponse(c, errno.ErrToken, "error, session is nil")
	}

	id, _ := c.GetQuery("id")
	codeId, _ := strconv.Atoi(id)

	au, err := sqlite.GetQrcode(codeId)
	if err != nil {
		handler.SendResponse(c, errno.ErrQuery, "error, list auth user!")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errno.SUCCESS,
		"data": au.Qrcode,
	})

}
