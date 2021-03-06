package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/jamiechapmanbrn/testmicro/handler"
	"github.com/jamiechapmanbrn/testmicro/subscriber"

	example "github.com/jamiechapmanbrn/testmicro/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.testmicro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.testmicro", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.testmicro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
