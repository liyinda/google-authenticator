package passport

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liyinda/google-authenticator/utils/errno"
	"github.com/liyinda/google-authenticator/utils/handler"
	"github.com/liyinda/google-authenticator/utils/jwt"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		handler.SendResponse(c, errno.ErrStruct, nil)
		return
	}

	if user.Username == "admin" && user.Password == "111111" {
		//将username存储到session中
		session := sessions.Default(c)
		session.Set("user", user.Username)
		err := session.Save()
		user_session := session.Get("user")
		log.Println("session", session, "user_session", user_session)
		if err != nil {
			handler.SendResponse(c, errno.ErrSession, nil)
			return
		}
		//生成用户token信息
		token, err := jwt.GenerateToken(user.Username, user.Password)
		if err != nil {
			handler.SendResponse(c, errno.ErrToken, nil)
		}
		data := gin.H{
			"token": token,
		}
		c.JSON(http.StatusOK, gin.H{
			"code": errno.SUCCESS,
			"data": data,
		})
		return
	}

	handler.SendResponse(c, errno.ErrNotLogin, "not login")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}
