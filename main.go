package main

import (
	"fmt"

	"example.com/PROJECT_1/db"
	v1 "example.com/PROJECT_1/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	db, err := db.DBconnection()
	fmt.Println(err)
	e := echo.New()

	conn, err := grpc.Dial(":2101", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(20000648), grpc.MaxCallSendMsgSize(20000648)))
	defer conn.Close()

	v1.SetUpRouter(e, db, conn)
	//fmt.Println("insert successfull")

	e.Logger.Fatal(e.Start(":4000"))

}
