package main

import (
    qrcode "github.com/yeqown/go-qrcode"
    //qrcode "github.com/skip2/go-qrcode"
    "fmt"
)

func main() {
   // var png []byte
   // png, _ = qrcode.Encode("https://example.org", qrcode.Medium, 256)
   // fmt.Println(png)

    qrc, err := qrcode.New("https://github.com/yeqown/go-qrcode")
    if err != nil {
        fmt.Printf("could not generate QRCode: %v", err)
    }

}
