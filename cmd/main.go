package main

import (
	"time"

	"github.com/JIeeiroSst/go-app/config"
	serive "github.com/JIeeiroSst/go-app/delivery/grpc"
	"github.com/JIeeiroSst/go-app/delivery/http"
	"github.com/JIeeiroSst/go-app/log"
	"github.com/JIeeiroSst/go-app/proto"
	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"google.golang.org/grpc"
	"github.com/labstack/echo/v4"
)
func main() {
	e:=echo.New()
	conn, err := grpc.Dial(config.Config.PORT, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	grpcClient:=proto.NewUserProfileClient(conn)
	defer func() {
		err := conn.Close()
		if err != nil {
			log.InitZapLog().Error("server running failed")
		}
	}()
	timeoutContext := time.Second
	repo:=mysql.NewMysqlConn(&config.Config.MysqlConfig)
	seriveGprc:=serive.NewService(repo,grpcClient,timeoutContext)
	http.NewHandler(e,seriveGprc)
	e.Logger.Fatal(e.Start(":8080"))
}