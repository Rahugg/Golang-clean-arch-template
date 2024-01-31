package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rahugg/Golang-clean-arch-template/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	config          *config.Config
	logger          *zap.SugaredLogger
	serviceProvider *serviceProvider
}

func NewApp(ctx *gin.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

func (a *App) initDeps(ctx *gin.Context) error {
	inits := []func(*gin.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initLogger,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ *gin.Context) error {
	config, err := config.New("config/user.yml")
	if err != nil {
		return err
	}

	a.config = &config

	return nil
}

func (a *App) initServiceProvider(_ *gin.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initLogger(_ *gin.Context) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	a.logger = logger.Sugar()

	return nil
}

func (a *App) initHTTPServer(_ *gin.Context) error {
	ctx, cancel := context.WithCancel(context.TODO())
	_ = ctx

	httpCfg := a.config.HttpServer
	server, err := http.NewServer(httpCfg.Port, httpCfg.ShutdownTimeout, router, a.logger, endpointHandler)
	if err != nil {
		a.logger.Panicf("failed to create server err: %v", err)
	}

	go func() {
		if err := server.Run(); err != nil {
			a.logger.Panicf("failed to run server err: %v", err)
		}
	}()

	defer func() {
		if err := server.Stop(); err != nil {
			a.logger.Panicf("failed close server err: %v", err)
		}
		a.logger.Info("server closed")
	}()

	a.gracefulShutdown(cancel)

	return nil
}

func (a *App) gracefulShutdown(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch
	signal.Stop(ch)
	cancel()
}
