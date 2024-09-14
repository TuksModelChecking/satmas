package main

import (
	"context"
	"fmt"
	"ui/go/api"
	"ui/go/codegen/ui/proto"
	"ui/go/experiment"
	"ui/go/mra"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// App struct
type App struct {
	ctx        context.Context
	logger     logger.Logger
	grpcServer api.GRPCServer
	*experiment.ExperimentStateControllerTSAdaptor
	experiment.ExperimentMetadataReader
	experiment.ExperimentResultReader
	experiment.ExperimentReader
}

func NewApp(
	logger logger.Logger,
	grpcServer api.GRPCServer,
	experimentStateController *experiment.ExperimentStateControllerTSAdaptor,
	experimentMetadataReader experiment.ExperimentMetadataReader,
	experimentResultReader experiment.ExperimentResultReader,
	experimentReader experiment.ExperimentReader,
) *App {
	return &App{
		logger:                             logger,
		grpcServer:                         grpcServer,
		ExperimentStateControllerTSAdaptor: experimentStateController,
		ExperimentMetadataReader:           experimentMetadataReader,
		ExperimentResultReader:             experimentResultReader,
		ExperimentReader:                   experimentReader,
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.ExperimentStateControllerTSAdaptor.SetAppContext(ctx)

	if err := a.ExperimentMetadataReader.InitialiseReader(); err != nil {
		a.logger.Fatal(
			fmt.Sprintf("error could not initialise experiment metadata reader: %s", err),
		)
	}

	// start grpc server in separate thread (blocking operation)
	go func() {
		if err := a.grpcServer.StartServer(); err != nil {
			a.logger.Fatal(fmt.Sprintf("error starting app grpc server: %s", err))
		}
	}()
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if err := a.ExperimentMetadataReader.CloseReader(); err != nil {
		a.logger.Fatal(fmt.Sprintf("error closing experiment metadata reader: %s", err))
	}
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	if err := a.grpcServer.StopServer(); err != nil {
		a.logger.Fatal(fmt.Sprintf("error stopping app grpc server: %s", err))
	}
	if err := a.ExperimentMetadataReader.CloseReader(); err != nil {
		a.logger.Fatal(fmt.Sprintf("error closing experiment metadata reader: %s", err))
	}
}

// utility functions
func (a *App) GenerateMRA(agents []*proto.Agent, resources []string) *mra.MRA {
	return mra.NewMRA(
		agents,
		resources,
	)
}
