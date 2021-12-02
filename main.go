package main

import (
	"blog/bootstrap"
	"blog/handler"
	blog "blog/proto/blog"
	limit "github.com/juju/ratelimit"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit/v2"
	"time"
)

func main() {
	bootstrap.Initialization()
	consul := consul.NewRegistry()
	rateLimitConfig := limit.NewBucketWithRate(float64(500), int64(1000))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.blog"),
		micro.Version("latest"),
		micro.Registry(consul),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(rateLimitConfig,false)),
	)
	service.Init()
	// Register Handler
	blog.RegisterBlogHandler(service.Server(), new(handler.Blog))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
