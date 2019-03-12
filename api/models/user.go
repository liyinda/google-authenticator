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

//func (Admin) TableName() string {
//    return "authcms_admin"
//}


//var GetPassword []Adminx
//var Authcms_admin []Admins

//获取管理员密码
func (authcms_admins *Authcms_admin) GetPassword(login_name string) (authcms_admin []Authcms_admin, err error) {
    if err = orm.Eloquent.Model(&authcms_admin).Where("login_name=?",login_name).First(&authcms_admin).Error; err != nil {
        return 

    }
    return 
}

//获取管理员列表
//func (admin *Admin) Admins() (admins []Admin, err error) {
//    if err = orm.Eloquent.Find(&admins).Error; err != nil {
//        return
//    }
//    return
//}



//type Userz struct {
//    ID       int64  `json:"id"`       // 列名为 `id`
//    Username string `json:"username"` // 列名为 `username`
//    Password string `json:"password"` // 列名为 `password`
//    Tassword string `json:"tassword"` // 列名为 `tassword`
//}
//type Userlist struct {
//    ID       int64  `json:"id"`       // 列名为 `id`
//    Username string `json:"username"` // 列名为 `username`
//    Password string `json:"password"` // 列名为 `password`
//    Tassword string `json:"tassword"` // 列名为 `tassword`
//}
//
//var Users []User

//列表
//func (user *User) Users() (users []User, err error) {
//func (user *Userlist) Userlist(id int64) (userz []Userlist, err error ) {
//    //id := c.Param("id")
//    //if err = orm.Eloquent.Find(&users).Error; err != nil {
//    if err = orm.Eloquent.Model(&userz).Where("id=?",id).First(&userz, id).Error; err != nil {
//        return 
//    }
//    return 
//}


//func (user *User) Password(id int64) (passwordUser User, err error) {
//    if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
//        return
//    }
//
//}

//修改
//func (user *User) Update(id int64) (updateUser User, err error) {
//
//    if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
//        return
//    }
//
//    //参数1:是要修改的数据
//    //参数2:是修改的数据
//    if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
//        return
//    }
//    return
//}
//
////查询
//
//func (user *User) Where(id int64) ( whereUser User, err error) {
//    //err := orm.Eloquent.Model(&Like{}).Where(&Like{Ip: ip, Ua: ua, Title: title}).Count(&count).Error
//    if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
//        return
//    }
//
//    whereUser = *user
//
//    return
//}
