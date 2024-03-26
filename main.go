package main

import (
	"context"
	"fmt"

	"github.com/hasura/ndc-sdk-go-reference/configuration"

	"github.com/hasura/ndc-sdk-go/connector"
)

func main() {
	var cli CLI
	if err := connector.StartCustom[Configuration, configuration.State](
		&cli,
		&Connector{},
		connector.WithMetricsPrefix("ndc_ref"),
		connector.WithDefaultServiceName("ndc_ref"),
	); err != nil {
		panic(err)
	}
}

type CLI struct {
	connector.ServeCLI
	Version struct{} `cmd:"" help:"Print the version."`
	Update  struct{} `cmd:"" help:"Update the configurations."`
}

func (cli *CLI) Execute(ctx context.Context, command string) error {
	logger := connector.GetLogger(ctx)
	switch command {
	case "version":
		logger.Info().Msg("0.1.0")
		return nil
	case "update":
		ctx := configuration.ConnectWithElasticsearch(context.Background())
		err := configuration.UpdatetConfigurations(ctx)
		if err == nil {
			logger.Info().Msg("Configuration Updated Succesfully.")
		}
		return err
	default:
		return fmt.Errorf("unknown command <%s>", command)
	}
}
