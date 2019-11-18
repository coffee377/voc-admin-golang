module github.com/coffee377/voc-admin

go 1.13

replace golang.org/x/net => github.com/golang/net v0.0.0-20190415214537-1da14a5a36f2

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190411191339-88737f569e3a

replace golang.org/x/text => github.com/golang/text v0.3.0

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190415145633-3fd5a3612ccd

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190415205951-2e9de471ebd3

require (
	github.com/astaxie/beego v1.12.0
	github.com/casbin/casbin v1.9.1
	github.com/casbin/casbin/v2 v2.0.1
	github.com/casbin/xorm-adapter v0.0.0-20191026011336-f11f9b23cc3d
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/wenzhenxi/gorsa v0.0.0-20190415061753-09b86265c6df // indirect
)
