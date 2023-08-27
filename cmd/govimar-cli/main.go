/*
Copyright © 2023 cheina97
*/
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/cheina97/govimar/cmd/govimar-cli/cmd"
	"github.com/cheina97/govimar/pkg/config"
)

func main() {
	cfg := config.GetConfig()
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	root := cmd.NewRootCmd(cfg)
	if err := root.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
