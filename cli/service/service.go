package service

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tedkahwaji/gcp-ints-cli/cli/googlecloud"
)

type Service struct {
	ProjectsClient googlecloud.ProjectsClient
}

func (s *Service) SearchProjects(cmd *cobra.Command, _ []string) {
	ctx := cmd.Context()
	query := "state:active AND NOT projectId=sys*"

	fmt.Println("Searching projects....")
	for result := range s.ProjectsClient.SearchProjects(ctx, query) {
		result, err := result.Item, result.Error
		if err != nil {
			fmt.Println("got an error ", err)
			return
		}
		fmt.Println(fmt.Sprintf("Project ID: %s, Name: %s, Parent: %s", result.ProjectId, result.Name, result.Parent))
	}

	fmt.Println("Done searching projects")

	fmt.Println("Searching folders....")
	for result := range s.ProjectsClient.SearchFolders(ctx, query) {
		result, err := result.Item, result.Error
		if err != nil {
			fmt.Println("got an error ", err)
			return
		}
		fmt.Println(fmt.Sprintf("Folder Name: %s, Display Name: %s, Parent: %s", result.Name, result.DisplayName, result.Parent))
	}
	fmt.Println("Done searching folders")
}
