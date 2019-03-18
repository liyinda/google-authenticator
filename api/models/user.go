package models

import (
    orm "github.com/liyinda/google-authenticator/api/database"
    "fmt"
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

//更新用户信息
func (authcms_users *Authcms_user) Useredit(id int64) (updateUser Authcms_user, err error) {
    //更新数据
    if err = orm.Eloquent.Select([]string{"id", "user_name"}).First(&updateUser, id).Error; err != nil {
        return
    }
    //参数1:是要修改的数据
    //参数2:是修改的数据
    if err = orm.Eloquent.Model(&updateUser).Updates(&authcms_users).Error; err != nil {
        return
    }
    return
}

//获取用户列表
func (authcms_users *Authcms_user) Userlist(page int64, limit int64) (users []Authcms_user, err error) {
    //查找数据
    fmt.Println(&users)
    //if page > 0 && limit > 0 {
        //db := orm.Eloquent.Limt(limt).Offset((page - 1) * limit)
    //}
    if err = orm.Eloquent.Limit(limit).Offset((page - 1) * limit).Find(&users).Error; err != nil {
        return
    }
    return
}


