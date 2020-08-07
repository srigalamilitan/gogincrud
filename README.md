# gogincrud
learning golanng with gin , gorm and wire

## GOLANG
1. Build 
  go build
2. Running Golang 
 PORT=8080 DB_URL="root:password_DB@(IP_ADDRESS:PORT)/gogin?charset=utf8&parseTime=True&loc=Local" ./gogincrud
 
 ## Docker
 1. Build Images
  at folder gogincrud run script
  docker build -t gogincrud:1.0.0 .
 2. Run Docker 
 docker run -it -p 8080:8080 --name "go-gin-crud" -e PORT=8080 -e DB_URL="root:PASSWORD_DB@(IP_ADDRESSDB:PORT)/gogin?charset=utf8&parseTime=True&loc=Local" -d gogincrud:1.0.0
 
 
 
 
 thanks for 
 https://hellokoding.com/golang/rest-api-wire-gin-gorm-mysql/
  
