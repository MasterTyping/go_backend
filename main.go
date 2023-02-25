package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/mastertyping/go_backend/go/src/routes"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Starting the program")

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	config := &firebase.Config{ProjectID: "nuritopia"}
	firebase, err := firebase.NewApp(context.Background(), config, opt)

	firebase.Messaging(context.Background())

	if err != nil {
		fmt.Println("Error initializing app:", err)
	}

	r := gin.Default()
	routes.Test(r)

	r.Run(":3001")
}
