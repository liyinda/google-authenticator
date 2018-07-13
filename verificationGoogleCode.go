package main

import (
    "fmt"
    "github.com/liyinda/google-authenticator/returncode"
    "net/http"
    "github.com/garyburd/redigo/redis"
    "log"
    "flag"
    "os"
)

var (
    logFileName = flag.String("log", "authenticator.log", "Log file name")
    listenAddress = flag.String("http.address", ":8082", "Address on HTTP Listen .")
    redisAddress = flag.String("redis.address", "127.0.0.1:6379", "Address on Redis Server .")

)



func handler(w http.ResponseWriter, r *http.Request) {
    vars := r.URL.Query()

    issuser,ok := vars["issuser"]

    if !ok || len(issuser) < 1 {
        fmt.Printf("param[issuser] a does not exist\n")
        return
    } else {
        //通过issuser值获取google验证码
        //c, err := redis.Dial("tcp", "127.0.0.1:6379")
        c, err := redis.Dial("tcp", *redisAddress)
        if err != nil {
            fmt.Println("Connect to redis error", err)
            return
        }
        defer c.Close()

        logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
        if logErr != nil {
                fmt.Println("Fail to find", *logFile, "cServer start Failed")
                os.Exit(1)
        }
        log.SetOutput(logFile)
        log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)


        randkey, err := redis.String(c.Do("GET", issuser[0]))
        if err != nil {
            log.Printf("用户不存在:%v \n", issuser)
            fmt.Fprint(w, "error")
        } else {

            code := vars["code"]
            //fmt.Printf("提交验证码%d\n",code[0])
            if !ok || len(code) < 1 {
                fmt.Printf("param[code] a does not exist\n")
                fmt.Fprint(w, "error")
                return
            } else {
                //判断google验证码与code是否一致
                result := returncode.ReturnCode(randkey)
                //result需要与code类型一致
                if code[0] == fmt.Sprint(result) {
                    fmt.Fprint(w, "ok")
                    log.Printf("认证成功:%v \n", issuser)
                } else {
                    fmt.Fprint(w, "error")
                    log.Printf("认证失败:%v \n", issuser)
                }
            
            }
        
                }

    }

}


func main() {
    flag.Parse()
    http.HandleFunc("/", handler)
    //http.ListenAndServe(":8082", nil)
    http.ListenAndServe(*listenAddress, nil)

    fmt.Println("redis:", *redisAddress)
}
