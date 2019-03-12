package database

import (
    "github.com/jinzhu/gorm"
    "fmt"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var Eloquent *gorm.DB

func init() {
    var err error
    Eloquent, err = gorm.Open("mysql", "root:liyinda@tcp(127.0.0.1:3306)/authcms?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if Eloquent.Error != nil {
        fmt.Printf("database error %v", Eloquent.Error)
    }

    //在go中需要定义一个struct， struct的名字就是对应数据库中的表名，
    //注意gorm查找struct名对应数据库中的表名的时候会默认把你的struct中的大写字母转换为小写并加上's'

    Eloquent.SingularTable(true)
}

