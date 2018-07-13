package main

import qrcode "github.com/skip2/go-qrcode"
import (
    "os"
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
    "math/rand"
)

//生成32位随机序列
var (
    codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    codeLen = len(codes)
)

func RandNewStr(len int) string {
    data := make([]byte, len)
    rand.Seed(time.Now().UnixNano())

    for i := 0; i < len; i++ {
        idx := rand.Intn(codeLen)
        data[i] = byte(codes[idx])
    }

    return string(data)
}



func main() {
    //连接redis服务端
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()

    //fmt.Println(RandStringBytes(32))

    var secretId string
    secretId = RandNewStr(32)

    //判断issuer是否存在,不存在写入redis服务端
    randkey, err := redis.String(c.Do("GET", os.Args[1]))
    if err != nil {
        //_, err = c.Do("SET", os.Args[1], RandStringBytes(32))
        _, err = c.Do("SET", os.Args[1], secretId)
        if err != nil {
            fmt.Println("redis set failed:", err)
        }

    } else {
        fmt.Printf("%v 用户已经注册,秘钥 %v\n", os.Args[1], randkey)
    }


    url := "otpauth://totp/liyinda.com?secret=" + secretId + "&issuer=" + os.Args[1]
    error := qrcode.WriteFile(url, qrcode.Medium, 256, "jpg/" + os.Args[1] + ".png")
    if error != nil {
        fmt.Println("write error")
    }
}
