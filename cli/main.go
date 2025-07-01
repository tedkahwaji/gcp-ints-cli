/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"context"

	"github.com/tedkahwaji/gcp-ints-cli/cli/cmd"
	"github.com/tedkahwaji/gcp-ints-cli/cli/googlecloud"
	"github.com/tedkahwaji/gcp-ints-cli/cli/service"
)

func main() {
	ctx := context.Background()

	projectsClient, err := googlecloud.NewProjectsClient(ctx)
	if err != nil {
		return
	}

	service := &service.Service{
		ProjectsClient: projectsClient,
	}

	cmd.Execute(service)
}
