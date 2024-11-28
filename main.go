package main

import (
	"log"

	"github.com/Muhfikri12/shipping/infra"
	"github.com/Muhfikri12/shipping/router"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := router.NewRoutes(*ctx)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
