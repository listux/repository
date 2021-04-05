package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/listux/repository/db"
	"github.com/listux/repository/example"
)

func main() {
	ctx := context.Background()
	conn, _ := db.NewConnection(ctx) // this method will initialize all db connections
	logger := log.New(os.Stderr, "", 0)

	router := gin.New()

	group := router.Group("/v1", conn.Connect()) // connect will construct inside "conn" a pointer to the selected firestore
	{
		mhs := example.NewMyHandlerStruct(logger, conn)
		group.GET("/myHandler", mhs.MyHandler)
	}

	// ...
	router.Run()
}
