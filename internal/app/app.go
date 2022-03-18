package app

import (
	"shortener/config"
	"shortener/internal/routers"
	srv "shortener/internal/server"
	"shortener/internal/service"
	"shortener/pkg/logging"
	"shortener/providers/bitly"
)

func RunApp(cfg *config.Config) {
	providerSetup := bitly.NewUrlService(cfg.Provider.Token, cfg.Provider.GroupGuid, cfg.Provider.Domain)
	services := service.AppServices(service.Deps{ProviderService: providerSetup})
	handler := routers.AppHandler(services)
	server := srv.NewServer(cfg, handler.Init())
	err := server.Run()
	if err != nil {
		logging.EasyLogFatal("app", "failed to run server!", "", err)
		return
	}
}
