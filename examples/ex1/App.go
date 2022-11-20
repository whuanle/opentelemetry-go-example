package main

import (
	"context"
	"go.opentelemetry.io/otel"
	"io"
	"log"
)

const name = "App"

type App struct {
	r io.Reader
	l *log.Logger
}

// NewApp returns a new App.
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r: r, l: l}
}

func (a *App) Run(ctx context.Context) {
	newCtx, span := otel.Tracer(name).Start(ctx, "Run")
	defer span.End()

	a.Run1(newCtx)
	a.Run2(newCtx)
}

func (a *App) Run1(ctx context.Context) {
	newCtx, span := otel.Tracer(name).Start(ctx, "Run1")
	defer span.End()

	a.Run1_1(newCtx)
}

func (a *App) Run1_1(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, "Run1_1")
	defer span.End()
}

func (a *App) Run2(ctx context.Context) {
	_, span := otel.Tracer(name).Start(ctx, "Run2")
	defer span.End()
}
