package util

import qrcode "github.com/skip2/go-qrcode"
import (
    "fmt"
    "time"
    "math/rand"
    "encoding/base64"
    "io/ioutil"
     "os"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

//生成32位随机序列
var (
    codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    codeLen = len(codes)    
    coder = base64.NewEncoding(base64Table)
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

func Base64Encode(encode_byte []byte) []byte {
    return []byte(coder.EncodeToString(encode_byte))
}

//创建二维码图片并返回base64编码字符串
func CreateQrcode(username string) (string, string, error) {
    var secretId string
    secretId = RandNewStr(32)

    url := "otpauth://totp/liyinda.com?secret=" + secretId + "&issuer=" + username
    qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode_jpg/" + username + ".png")
    //将存入的图片转换为base64格式
    file, err := os.Open("qrcode_jpg/" + username + ".png")
    if err != nil {
        fmt.Println("无法打开二维码图片")
    }
    data, err := ioutil.ReadAll(file)
    return secretId, "data:image/png;base64," + string(Base64Encode(data)), err

}

