package service

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tedkahwaji/gcp-ints-cli/cli/googlecloud"
)

type Service struct {
	ProjectsClient googlecloud.ProjectsClient
}

func (s *Service) SearchProjects(cmd *cobra.Command, args []string) {
	for result := range s.ProjectsClient.SearchProjects(cmd.Context(), "state:active AND NOT projectId=sys*") {
		result, err := result.Item, result.Error
		if err != nil {
			fmt.Println("got an error ", err)
			return
		}
		fmt.Println(result.ProjectId)
	}
}
