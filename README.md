### INTRODUCE
 * Using language golang with echo framework and beego orm(connect database management system mysql and create model define) to build api login
 * Encrypt password generate token to user authentication and authorization for system.
### Instal lib
	go mod init github.com/auth
	go get github.com/astaxie/beego v1.12.3 // indirect
	go get github.com/beego/beego/v2 v2.0.3 // indirect
	go get github.com/go-sql-driver/mysql v1.6.0 // indirect
	go get github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	go get github.com/golang/glog v1.0.0 // indirect
	go get github.com/labstack/echo/v4 v4.7.2 // indirect
	go get golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
>Using command "go run .\main.go" to build system.

#### local.env sample
```
DATABASE_TYPE=mysql
USER_DB=root
PASS_WORD_DB=1234
SCHEMA_DB=go_auth
```
