package main

import (
	"blog/bootstrap"
	"blog/handler"
	blog "blog/proto/blog"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	bootstrap.Initialization()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.blog"),
		micro.Version("latest"),
	)
	// Register Handler
	blog.RegisterBlogHandler(service.Server(), new(handler.Blog))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
