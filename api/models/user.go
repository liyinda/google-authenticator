package models

import (
    orm "github.com/liyinda/google-authenticator/api/database"
)

// 用户登录 from JSON
type LoginJson struct {
        Loginname     string `form:"loginname" json:"loginname" xml:"loginname"  binding:"required"`
        Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//管理员数据表
type Authcms_admin struct {
    ID       int64  `json:"id"`       // 列名为 `id`
    Login_name string `json:"login_name"` // 列名为 `login_name`
    Real_name string `json:"real_name"` // 列名为 `username`
    Password string `json:"password"` // 列名为 `password`
    Phone string `json:"phone"` // 列名为 `phone`
    Email string `json:"email"` // 列名为 `email`
    Salt string `json:"salt"` // 列名为 `salt`
    Last_login int64 `json:"last_login"`       // 列名为 `last_login`
    Last_ip string `json:"last_ip"` // 列名为 `last_ip`
    Status int64 `json:"status"` // 列名为 `status`
    Create_id int64 `json:"create_id"` // 列名为 `create_id`
    Update_id int64 `json:"update_id"` // 列名为 `update_id`
    Create_time int64 `json:"create_time"` // 列名为 `create_time`
    Update_time int64 `json:"Update_time"` // 列名为 `update_time`
}

//获取管理员密码
func (authcms_admins *Authcms_admin) GetPassword(login_name string) (authcms_admin []Authcms_admin, err error) {
    if err = orm.Eloquent.Model(&authcms_admin).Where("login_name=?",login_name).First(&authcms_admin).Error; err != nil {
        return 

    }
    return 
}

