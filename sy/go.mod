module sy

go 1.15

require (
	fx v0.0.0-00010101000000-000000000000
	github.com/auyer/steganography v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-basic/uuid v1.0.0
	github.com/go-ini/ini v1.62.0
	github.com/spf13/cast v1.3.1
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201012173705-84dcc777aaee
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

 
replace fx => github.com/panyuan5056/fx v0.4.0 // indirect
 