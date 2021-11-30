package main

import (
	"blog/bootstrap"
	"blog/handler"
	blog "blog/proto/blog"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	bootstrap.Initialization()
	consul := consul.NewRegistry()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.blog"),
		micro.Version("latest"),
		micro.Registry(consul),
	)
	// Register Handler
	blog.RegisterBlogHandler(service.Server(), new(handler.Blog))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
