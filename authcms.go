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
    //"github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"
    "github.com/gin-contrib/cors"
    "github.com/liyinda/google-authenticator/api/router"
    //orm "github.com/liyinda/google-authenticator/api/database"
    //"fmt"
)

func main() {
    router := router.InitRouter()
    //服务器session
    store := sessions.NewCookieStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))

    //容许跨域访问
    //vue-admin需要单独添加("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") 
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
        AllowHeaders:     []string{"Content-Type,Authorization,X-Token"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "*"
        },
    }))

    router.Run(":8888")
}
