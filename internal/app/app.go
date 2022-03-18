package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shortener/config"
	"shortener/internal/routers"
	srv "shortener/internal/server"
	"shortener/internal/service"
	"shortener/pkg/logging"
	"shortener/providers/bitly"
	"syscall"
	"time"
)

func RunApp(cfg *config.Config) {
	providerSetup := bitly.NewUrlService(cfg.Provider.Token, cfg.Provider.URL, cfg.Provider.Domain)
	services := service.AppServices(service.Deps{ProviderService: providerSetup})
	handler := routers.AppHandler(services)
	server := srv.NewServer(cfg, handler.Init())
	logging.EasyLogInfo("app", "have a nice day! server started at: ", fmt.Sprintf("%s:%s",
		cfg.HTTP.Host, cfg.HTTP.Port))
	err := server.Run()
	if err != nil {
		logging.EasyLogFatal("app", "failed to run server!", "", err)
		return
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := server.Stop(ctx); err != nil {
		logging.EasyLogError("app", "failed to safe shutdown server", "", err)
	}
}
