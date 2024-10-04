package docker

import (
	"context"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type DockerImage struct {
	ID      string
	Tag     string
	Created int64
}

type DockerWrapper struct {
	client *client.Client
}

func (c *DockerWrapper) NewClient() {
	var err error
	c.client, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err.Error())
	}
}

func (c *DockerWrapper) CloseClient() {
	c.client.Close()
}

func (c *DockerWrapper) GetImages() []DockerImage {
	images, err := c.client.ImageList(context.Background(), image.ListOptions{All: true})

	if err != nil {
		panic(err)
	}

	var dockerImages []DockerImage
	for _, img := range images {
		if len(img.RepoTags) == 0 {
			dockerImages = append(dockerImages, DockerImage{
				ID:      img.ID,
				Tag:     "<none>:<none>",
				Created: img.Created,
			})
		} else {
			for _, tag := range img.RepoTags {
				dockerImages = append(dockerImages, DockerImage{
					ID:      img.ID,
					Tag:     tag,
					Created: img.Created,
				})
			}
		}
	}

	return dockerImages
}
