/*
Copyright 2019 The Google Authenticator Authors.
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
    //"github.com/gin-gonic/contrib/sessions"
    //"github.com/gin-contrib/cors"
    "github.com/liyinda/google-authenticator/api/router"
    //orm "github.com/liyinda/google-authenticator/api/database"
)

func main() {
    router := router.InitRouter()
    //服务器session
    //store := sessions.NewCookieStore([]byte("secret"))
    //router.Use(sessions.Sessions("mysession", store))

    //容许跨域访问
    //vue-admin需要单独添加("Access-Control-Allow-Headers", "Content-Type,Authorization,X-Token") 
    //router.Use(Cors())
//    router.Use(cors.New(cors.Config{
//        AllowOrigins:     []string{"http://192.168.30.18"},
//        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
//        AllowHeaders:     []string{"XMLHttpRequest","Content-Type","Authorization","X-Token","Origin","Access-Control-Allow-Origin",},
//        //AllowHeaders:     []string{"XMLHttpRequest,Content-Type,Authorization,X-Token,Origin,Access-Control-Allow-Origin"},
//        ExposeHeaders:    []string{"Content-Length"},
//        AllowCredentials: true,
//        AllowOriginFunc: func(origin string) bool {
//            return true
//            //return origin == "http://192.168.30.18"
//            //return origin == "http://101.200.42.56:8888"
//        },
//    }))

    router.Run(":8888")
}

