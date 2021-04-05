package main

import (
	"boundedinfinity/echo_blueprint/service"
	"boundedinfinity/echo_blueprint/web"
)

func main() {
	service := service.New()
	web := web.New(service)

	if err := web.Serve(); err != nil {
		panic(err)
	}
}
