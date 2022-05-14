package main

import (
	"github.com/rewle/service-select-participants/internal/config"
	"github.com/rewle/service-select-participants/internal/db"
	"github.com/rewle/service-select-participants/internal/logger"
	"github.com/rewle/service-select-participants/internal/participant"
	"github.com/rewle/service-select-participants/internal/server"
	"github.com/rewle/service-select-participants/internal/utils"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()

	utils.PanicOnErr(container.Provide(config.Init))
	utils.PanicOnErr(container.Provide(logger.Init))
	utils.PanicOnErr(container.Provide(participant.Init))
	utils.PanicOnErr(container.Provide(server.Init))
	utils.PanicOnErr(container.Provide(db.Init))

	return container
}

func main() {
	container := buildContainer()

	utils.PanicOnErr(container.Invoke(func(server *server.Server) {
		server.Run()
	}))
}
