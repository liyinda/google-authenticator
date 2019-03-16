package models

import (
    orm "github.com/liyinda/google-authenticator/api/database"
)


//用户数据表
type Authcms_user struct {
    ID       int64  `json:"id"`       // 列名为 `id`
    User_name string `json:"user_name"` // 列名为 `login_name`
    Real_name string `json:"real_name"` // 列名为 `username`
    Password string `json:"password"` // 列名为 `password`
    Phone string `json:"phone"` // 列名为 `phone`
    Email string `json:"email"` // 列名为 `email`
    Secretid string `json:"secretid"` // 列名为 `secretid`
    Salt string `json:"salt"` // 列名为 `salt`
    Last_login int64 `json:"last_login"`       // 列名为 `last_login`
    Last_ip string `json:"last_ip"` // 列名为 `last_ip`
    Status int64 `json:"status"` // 列名为 `status`
    Create_id int64 `json:"create_id"` // 列名为 `create_id`
    Update_id int64 `json:"update_id"` // 列名为 `update_id`
    Create_time int64 `json:"create_time"` // 列名为 `create_time`
    Update_time int64 `json:"update_time"` // 列名为 `update_time`
    Qrcode string `json:"qrcode"` // 列名为 `update_time`
}

//创建新用户
func (authcms_users *Authcms_user) Useradd() (id int64, err error) {
    //添加数据
    result := orm.Eloquent.Create(&authcms_users)
    id = authcms_users.ID
    if result.Error != nil {
        err = result.Error
        return
    }
    return

}
