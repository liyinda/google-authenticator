package models

import (
     // orm "github.com/liyinda/google-authenticator/api/database"

)

// 用户登录 from JSON
type LoginJson struct {
        Loginname     string `form:"loginname" json:"loginname" xml:"loginname"  binding:"required"`
        Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

