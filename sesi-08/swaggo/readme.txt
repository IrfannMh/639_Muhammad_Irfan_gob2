

step swago:
mkdir swaggo
cd swaggo
go mod init swaggo

generate swaggo docs:
swag init -g (nama file.go)

library yang diinstall:
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template

go get -u github.com/gorilla/mux

global :
go get -u github.com/swaggo/swag/cmd/swag@1.8.6
