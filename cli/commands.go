package main

import (
	"context"

	"github.com/hashicorp/consul-k8s/cli/cmd/common"
	"github.com/hashicorp/consul-k8s/cli/cmd/install"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/cli"
)

func InitializeCommands(ctx context.Context, log hclog.Logger) (*common.BaseCommand, map[string]cli.CommandFactory) {

	baseCommand := &common.BaseCommand{
		Ctx: ctx,
		Log: log,
	}

	commands := map[string]cli.CommandFactory{
		"install": func() (cli.Command, error) {
			return &install.Command{
				BaseCommand: baseCommand,
			}, nil
		},
	}

	return baseCommand, commands
}
