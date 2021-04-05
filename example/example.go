package example

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/listux/repository/db"
)

type MyHandlerStruct struct {
	*log.Logger
	*db.Connection
}

func NewMyHandlerStruct(log *log.Logger, conn *db.Connection) *MyHandlerStruct {
	return &MyHandlerStruct{
		log,
		conn,
	}
}

func (t *MyHandlerStruct) MyHandler(ctx *gin.Context) {
	// so instead of the original suggestion:
	// t.FirestoreClient.Collection("customers")...
	t.Firestore(ctx).Collection("customers").Doc("aaaa").Get(ctx)
}
