/*
Copyright 2018 The AmrToMp3 Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/


package main

import (
    _ "fmt"
    "github.com/gin-gonic/gin"
    //"math/rand"
    "net/http"
)

type LoginForm struct {
    User     string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
}


func main() {
    router := gin.Default()
    passport := router.Group("/passport")
    {
        passport.POST("/login", func(c *gin.Context) {
            // you can bind multipart form with explicit binding declaration:
            // c.ShouldBindWith(&form, binding.Form)
            // or you can simply use autobinding with ShouldBind method:
            var form LoginForm
            if c.ShouldBind(&form) == nil {
                if form.User == "user" && form.Password == "password" {
                    c.JSON(200, gin.H{"status": "you are logged in"})
                } else {
                    c.JSON(401, gin.H{"status": "unauthorized"})
                }
            }
        })
    }



    //定义默认路由
    router.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{
            "status": 404,
            "error":  "404, page not exists!",
        })
    })
    router.Run(":8888")
}
