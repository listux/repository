package db

import (
	"context"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
)

const (
	firestoreKey = "firestore"
)

type Connection struct {
	demoFS *firestore.Client // UNEXPORTED
	prodFS *firestore.Client // UNEXPORTED
}

// NewConnection initializes a base connection - currently Firestore exported field is nil.
func NewConnection(ctx context.Context) (*Connection, error) {
	demo, _ := firestore.NewClient(ctx, "demo")
	prod, _ := firestore.NewClient(ctx, "prod")
	return &Connection{
		demoFS: demo,
		prodFS: prod,
	}, nil
}

// Connect will store inside connection, a pointer to a firestore connection based on mode.
func (c *Connection) Connect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetBool("demoMode") {
			ctx.Set(firestoreKey, c.demoFS)
		} else {
			ctx.Set(firestoreKey, c.prodFS)
		}
	}
}

// Firestore returns a pointer to the selected firestore depending on firestoreKey on context
// (not related to gin.Context)
func (c *Connection) Firestore(ctx context.Context) *firestore.Client {
	return ctx.Value(firestoreKey).(*firestore.Client)
}
