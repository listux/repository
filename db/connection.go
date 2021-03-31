package db

import (
	"context"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
)

type Connection struct {
	Firestore *firestore.Client
	demoFS    *firestore.Client // UNEXPORTED
	prodFS    *firestore.Client // UNEXPORTED
}

func NewConnection(ctx context.Context) (*Connection, error) {
	demo, _ := firestore.NewClient(ctx, "demo")
	prod, _ := firestore.NewClient(ctx, "prod")
	return &Connection{
		demoFS: demo,
		prodFS: prod,
	}, nil
}

// Connect will store inside firestore, a pointer to a connection based on mode.
func (c *Connection) Connect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetBool("demoMode") {
			c.Firestore = c.demoFS
		} else {
			c.Firestore = c.prodFS
		}
	}
}
