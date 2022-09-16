module google-authenticator

go 1.15

require (
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ini/ini v1.66.6
	github.com/go-xorm/xorm v0.7.9
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/lib/pq v1.1.1
	github.com/liyinda/google-authenticator v0.0.0-20220104094512-245494761dad
	github.com/mattn/go-sqlite3 v1.14.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
)

replace github.com/liyinda/google-authenticator => ./
