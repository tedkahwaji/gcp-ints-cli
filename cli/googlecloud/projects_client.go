package googlecloud

import (
	"context"
	"iter"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
)

//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE

type ProjectsClient interface {
	GetProject(ctx context.Context, projectID string) (*resourcemanagerpb.Project, error)
	SearchProjects(ctx context.Context, query string) iter.Seq[APIResult[*resourcemanagerpb.Project]]
	Close() error
}

func NewProjectsClient(ctx context.Context) (ProjectsClient, error) {
	client, err := resourcemanager.NewProjectsClient(
		ctx,
		option.WithTelemetryDisabled(),
	)
	if err != nil {
		return nil, err
	}

	return &projectsClient{client: &googleProjectsClient{client: client}}, nil
}

type projectsClient struct {
	client GoogleProjectsClient
}

func (c *projectsClient) GetProject(ctx context.Context, projectID string) (*resourcemanagerpb.Project, error) {
	project, err := c.client.GetProject(ctx, &resourcemanagerpb.GetProjectRequest{Name: "projects/" + projectID})
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (c *projectsClient) SearchProjects(ctx context.Context, query string) iter.Seq[APIResult[*resourcemanagerpb.Project]] {
	return buildIterSeq(c.client.SearchProjects(ctx, &resourcemanagerpb.SearchProjectsRequest{Query: query}))
}

func (c *projectsClient) Close() error {
	return c.client.Close()
}

// wrappers/interfaces for Google SDK

type GoogleProjectsClient interface {
	GetProject(ctx context.Context, req *resourcemanagerpb.GetProjectRequest, opts ...gax.CallOption) (*resourcemanagerpb.Project, error)
	SearchProjects(ctx context.Context, req *resourcemanagerpb.SearchProjectsRequest, opts ...gax.CallOption) Iterator[*resourcemanagerpb.Project]
	Close() error
}

var _ GoogleProjectsClient = (*googleProjectsClient)(nil)

type googleProjectsClient struct {
	client *resourcemanager.ProjectsClient
}

func (c *googleProjectsClient) GetProject(ctx context.Context, req *resourcemanagerpb.GetProjectRequest, opts ...gax.CallOption) (*resourcemanagerpb.Project, error) {
	return c.client.GetProject(ctx, req, opts...)
}

func (c *googleProjectsClient) SearchProjects(ctx context.Context, req *resourcemanagerpb.SearchProjectsRequest, opts ...gax.CallOption) Iterator[*resourcemanagerpb.Project] {
	return c.client.SearchProjects(ctx, req, opts...)
}

func (c *googleProjectsClient) Close() error {
	return c.client.Close()
}
