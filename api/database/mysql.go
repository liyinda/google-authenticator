package database

import (
    "github.com/jinzhu/gorm"
    "fmt"
)

var Eloquent *gorm.DB

func init() {
    var err error
    Eloquent, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/authcms?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if Eloquent.Error != nil {
        fmt.Printf("database error %v", Eloquent.Error)
    }
}

